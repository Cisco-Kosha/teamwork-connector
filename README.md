# Kosha Teamwork Connector

![Teamwork](images/teamwork.png)

Teamwork is a project management and collaboration platform that helps teams plan, manage, and deliver projects of any size.

The Kosha Teamwork connector enables you to perform REST API operations from the Teamwork API in your Kosha workflow or custom application. Using the Kosha Teamwork connector, you can directly access the Teamwork platform to:

* Manage people information
* Get milestones
* Manage project lifecycles, risks, and tasks
* Manage project risks

## Useful Actions 

Use the Kosha Teamwork connector to perform several useful operations to manage your Teamwork resources. Refer to the Kosha Teamwork connector [API specification](openapi.json) for details.

### People

Get all available people data for projects, milestones, and tasks.

### Projects

Retrieve all available projects within Teamwork and manage their life cycles (for resources that the authenticated user is associated with).

### Milestones

Get all milestones for on projects that the authenticated user is associated with. Use filters to return only those milestones that are incomplete, completed, upcoming, or late.

### Risks

Manage project risks.

### Tasks/Tasklist

Manage projects tasks and task lists.

## Authentication

To authenticate when provisioning the Kosha Teamwork connector, you need your:

* Teamwork URL
* API key.

See the [Teamwork API docs](https://apidocs.teamwork.com/docs/teamwork/d1b2de52c3cec-api-key-and-url) for details on retrieving your URL and API key. 
