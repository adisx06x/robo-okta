# Robo-Okta

## Introduction

Robo-Okta is a cli that automates cumbersome Okta tasks done in the UI.

## Installation

```console
brew install adisx06x/tap/robo-okta
```

## Initial Setup

To use Robo-Okta, you must have your Okta org URL and proper API credentials.

Robo-Okta, for now, only takes these values as environment variables, thus, you must export them on your terminal before using Robo-Okta.

1. ORG URL

    ```export OKTA_CLIENT_ORGURL=https://example.okta.com```

1. Okta API Token

    ```export OKTA_CLIENT_TOKEN=EnterTokenHere```

## Using Robo-Okta

### Examples

1. Get User

    ```robo-okta get-user oktausername@example.com```

### Add csv list of users to a group

To successfully run this command, make sure you have a csv of all the users (just the login) that you want to add to a group.

You must also specify the group ID of the user. There are many ways to get this. The easiest is to grab it from the URL of the group. It is the last string of numbers after /group/`groupidhere`

1. Add csv list of users to a group command

    ```robo-okta add-users-to-group -f filename.csv -g groupid```

### Under Construction
