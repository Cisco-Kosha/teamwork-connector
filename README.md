# Kosha Teamwork Connector

![Teamwork](images/teamwork.png)

Teamwork is a project management and collaboration platform that helps teams plan, manage, and deliver projects of any size.

Using the Kosha Teamwork connector, you can perform REST API operations such as reading, modifying, adding, or deleting data from your projects.

This Kosha Teamwork connector exposes REST API endpoints to perform any operations on Teamwork APIs in a simple, quick and intuitive fashion. Using the Teamwork API, your Kosha workflow or application can directly access the Teamwork platform to:

* Manage people information
* Get milestones
* Manage project lifecycle
* Manage project risks
* Manage projects tasks and task lists
* Manage time entry information

## Useful Actions 

Use the Kosha Teamwork connector to perform several useful operations to manage your Teamwork resources. Refer to the Kosha Teamwork connector [API specification](openapi.json) for details.

* People: Get all available people data for projects, milestones, and tasks.
* Projects: Retrieve all available projects within Teamwork and manage their life cycles (for resources that the authenticated user is associated with).
* Milestones: Get all milestones for on projects that the authenticated user is associated with. Use filters to return only those milestones that are incomplete, completed, upcoming, or late.
* Risks: Manage project risks.
* Tasks / Tasklist: Manage projects tasks and task lists.

## Example Usage

The following request creates a new project: 

```
curl --request POST \
  --url https://stoplight.io/mocks/teamwork-dot-com/teamwork/42258908/projects.json \
  --header 'Authorization: Basic aXNhYWM6' \
  --header 'Content-Type: application/json' \
  --data '{
  "project": {
    "name": "New Project",
    "description": "My Big New Project",
    "start-date": "20230131",
    "end-date": "20230131",
  }
```

## Authentication

To use the Kosha Teamwork Connector, you need your Teamwork URL and API Key.

See the [Teamwork API docs](https://apidocs.teamwork.com/docs/teamwork/d1b2de52c3cec-api-key-and-url) for details. 

