#!/usr/bin/env bash
# Generate Go client SDK from OpenAPI specification.

openapi_generator_tag="v5.1.1"

cloud_export_spec_filename="cloud_export.openapi.json"
cloud_export_spec_url="https://raw.githubusercontent.com/kentik/api-schema-public/master/gen/openapiv2/kentik/cloud_export/v202101beta1/cloud_export.swagger.json"
synthetics_spec_filename="synthetics.openapi.json"
synthetics_spec_url="https://raw.githubusercontent.com/kentik/api-schema-public/master/gen/openapiv2/kentik/synthetics/v202101beta1/synthetics.swagger.json"

function run() {
    check_prerequisites

    download_openapi_spec
    generate_cloud_export_client
    generate_synthetics_client
    go fmt ./...
}

function check_prerequisites() {
    stage "Checking prerequisites"

    if ! docker --version > /dev/null 2>&1; then
        echo "You need to install Docker to run the generator"
        exit 1
    fi

    echo "Done"
}

function download_openapi_spec() {
    wget "$cloud_export_spec_url" --output-document "$cloud_export_spec_filename"
    wget "$synthetics_spec_url" --output-document "$synthetics_spec_filename"
}

function generate_cloud_export_client() {
    cloudexport_package="cloudexport"
    cloudexport_client_output_dir="kentikapi/cloudexport"

    generate_go_client_from_openapi_spec "$cloud_export_spec_filename" "$cloudexport_package" "$cloudexport_client_output_dir"
    change_ownership_to_current_user "$cloudexport_client_output_dir"
    cleanup_non_needed_files "$cloudexport_client_output_dir"
}

function generate_synthetics_client() {
    synthetics_package="synthetics"
    synthetics_client_output_dir="kentikapi/synthetics"

    generate_go_client_from_openapi_spec "$synthetics_spec_filename" "$synthetics_package" "$synthetics_client_output_dir"
    change_ownership_to_current_user "$synthetics_client_output_dir"
    cleanup_non_needed_files "$synthetics_client_output_dir"
}

function generate_go_client_from_openapi_spec() {
    stage "Generating Go client from OpenAPI specification"

    spec_file="$1"
    package="$2"
    output_dir="$3"

    docker run --rm -v "$(pwd):/local" \
        "openapitools/openapi-generator-cli:$openapi_generator_tag" generate \
        -i "/local/$spec_file" \
        -g go \
        --additional-properties=enumClassPrefix=true \
        --package-name "$package" \
        -o "/local/$output_dir"

    echo "Done"
}

function change_ownership_to_current_user() {
    dir="$1"
    stage "Changing ownership of $dir to $USER:$USER"

    sudo chown -R "$USER:$USER" "$dir" # by default the generated output is in user:group root:root

    echo "Done"
}

function cleanup_non_needed_files() {
    stage "Removing non-needed generated files"

    generated_content_dir="$1"

    rm "$generated_content_dir/.travis.yml"
    rm "$generated_content_dir/git_push.sh"
    rm "$generated_content_dir/go.mod"
    rm "$generated_content_dir/go.sum"
    rm "$generated_content_dir/README.md"
    rm -rf "$generated_content_dir/api"
    rm -rf "$generated_content_dir/docs"

    echo "Done"
}

function stage() {
    BOLD_BLUE="\e[1m\e[34m"
    RESET="\e[0m"
    msg="$1"

    echo
    echo -e "$BOLD_BLUE$msg$RESET"
}

run
