# SCIM Setup

This application is still very much work in progress, but the idea is that it should be able to do:

- Read user emails from a file :white_check_mark:
- Get the users object ids from the corresponding Azure AD user :white_check_mark:
- Add the users to an Azure AD group :white_check_mark:
- Configure and Azure Enterprise Application :x:
- Configure SCIM with the Azure Enterprise Application :x:

## Prerequisites

### Environment file

Create a file _.scim-setup_ and either place it in the $HOME directory or in the current working 
directory.

```bash
AZURE_CLIENT_ID=<REDACTED>
AZURE_CLIENT_SECRET=<REDACTED>
AZURE_TENANT_ID=<REDACTED>
AZURE_GROUP_OBJECT_ID=<REDACTED>
```

The AZURE_GROUP_OBJECT_ID is the ObjectId of the Azure AD group you want to add users to.

The AZURE_CLIENT_ID, AZURE_CLIENT_SECRET and AZURE_TENANT_ID must be the credentials for a service 
principal that is an owner of the Azure AD Group defined through AZURE_GROUP_OBJECT_ID.

You do not need to source the environment file or export these environment variables manually. 
The application will read the values directly from the _.scim-setup_ file.

### Users file

The users you want to add to the Azure AD Group defined through AZURE_GROUP_OBJECT_ID should be in the 
format userid@domain.tld and be put one per line in a file called users.txt (or any other name you prefer).

You may put the file anywhere you like on disk, for instance in the /tmp directory.

## Running the program

```bash
./scim-setup /tmp/users.txt # Full path to the users.txt file
```

## Development practices

WIP

### Build instructions

WIP

### Test instructions

WIP