# awsls

A list command for AWS resources.

[![Release](https://img.shields.io/github/release/jckuester/awsls.svg?style=for-the-badge)](https://github.com/jckuester/awsls/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE.md)
[![Travis](https://img.shields.io/travis/jckuester/awsls/master.svg?style=for-the-badge)](https://travis-ci.org/jckuester/awsls)

awsls supports listing of [over 200 types of resources](#supported-resources)
across {{ len .Services }} different AWS services. The goal is to code-generate a list function for
every AWS resource that is covered by the Terraform AWS Provider (currently over 500). If you want to contribute,
[the generator is here](./gen).

If you encounter any issue with `awsls` or have a feature request, 
please open an issue or write me on [Twitter](https://twitter.com/jckuester).

Happy listing!

**Note:** If you're also looking for an easy but powerful way to delete AWS resources, pipe the output of `awsls` into its new sibling
[`awsrm`](https://github.com/jckuester/awsrm) via Unix-pipes and use well-known standard tooling such as `grep` for filtering.

## Features

* List multiple types of resources at once by using glob patterns
  (e.g., `"aws_iam_*"` lists all IAM resources and `"*"` all resources in your account)
* **New:** List resources across multiple accounts and regions by using the `--profiles` and `--regions` flag
  (e.g., `-p account1,account2 -r us-east-1,us-west-2`)
* Show any resource attribute documented in the [Terraform AWS Provider](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)
  (e.g., `-a private_ip,tags` lists the IP and tags for resources of type [`aws_instance`](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance#attributes-reference))

## Examples

### List various resource attributes

Use Terraform resource types to tell `awsls` which resources to list. For example, `awsls aws_instance` shows
all EC2 instances. In addition to the default attributes `TYPE`, `ID`, `REGION`, and `CREATED` timestamp, additional attributes
can be displayed via the `--attributes <comma-separated list>` flag. Every attribute in the Terraform documentation 
([here](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance#attributes-reference) are the attributes for `aws_instance`) is a valid one:

![](img/instance.gif)

### List multiple resource types at once (via glob patterns)

For example, `awsls "aws_iam_*` lists all IAM resources:

![](img/iam.gif)

### List across multiple accounts and regions

To use specific profiles and/or regions, use the `-p (--profiles)` or `-r (--regions)` flags. For example,
`-p myaccount1,myaccount2 -r us-east-1,us-west-2` lists resources in every permutation of the given profiles and regions, 
i.e., resources in region `us-west-2` and `us-east-1` for account `myaccount1` as well as `myaccount2`:

![](img/multi-profiles-and-regions.gif)

## Usage

	awsls [flags] <resource_type glob pattern>

To see options available run `awsls --help`.

## Installation

### Binary Releases

You can download a specific version of awsls on the [releases page](https://github.com/jckuester/awsls/releases) or
install it the following way to `./bin/`:

```bash
curl -sSfL https://raw.githubusercontent.com/jckuester/awsls/master/install.sh | sh -s v0.6.1
```

### Homebrew

Homebrew users can install by:

```bash
brew install jckuester/tap/awsls
```

For more information on Homebrew taps please see the [tap documentation](https://docs.brew.sh/Taps).

## Credentials, profiles and regions

If the  `--profiles` and/or `--regions` flag is unset, `awsls` will follow the usual 
[AWS CLI precedence](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html#cli-configure-quickstart-precedence)
of first trying to find credentials, profiles and/or regions via [environment variables](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html),
and so on.

For example, if using `--profiles foo,bar`, but not setting the regions flag, 
`awsls` will first try to use the region from an environment variable (e.g., `AWS_DEFAULT_REGION`)
and second will try to use the default region for each profile from `~/.aws/config`.

The `--all-profiles` flag will use all profiles from `~/.aws/config`, or if `AWS_CONFIG_FILE=/my/config` is set, from
`/my/config` otherwise.

## Supported resources

Currently, all {{ .SupportedResourceTypeCount }} resource types across {{ len .Services }} services in the table below can be listed with awsls. The `Tags` column shows if a resource
supports displaying tags, the `Creation Time` column if a resource has a creation timestamp, and the `Owner` column if
resources are pre-filtered belonging to the account owner.

Note: the prefix `aws_` for resource types is now optional. This means, for example,
`awsls aws_instance` and `awsls instance` are both valid commands.

| Service / Type | Tags | Creation Time | Owner
| :------------- | :--: | :-----------: | :---:
{{ range $service := .Services }}| **{{ $service.Name }}** |
{{ range $rType := $service.TerraformResourceTypes -}}
| {{ $rType.Name }} | {{ if $rType.Tags }} x {{ end }} | {{ if $rType.CreationTime }} x {{ end }} |{{ if $rType.Owner }} x |{{ end }}
{{ end }}
{{- end }}