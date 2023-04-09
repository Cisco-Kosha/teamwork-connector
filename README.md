# Kosha Teamwork Connector

Teamwork is a project management platform that provides a variety of tools to keep teams on track.

The connector APIs allow you to perform 'RESTful' operations such as reading, modifying, adding or deleting data from your projects. The APIs also support Cross-Origin Resource Sharing (CORS).


![Teamwork](images/teamwork.jpg)


## What is possible with Teamwork APIs?

The Teamwork APIs provide your applications with direct access to the Teamwork, giving you the ability to:

* Manage people information
* Get milestones
* Manage project lifecycle
* Manage project risks
* Manage projects tasks and tasklists
* Manage timeentry information
and much more!


This Connector API exposes REST API endpoints to perform any operations on Teamwork APIs in a simple, quick and intuitive fashion.

It describes various API operations, related request and response structures, and error codes.

## Useful Actions 

### People

Get all available people data over projects, milestones, for a task etc.

### Projects

Retrieve all available projects within Teamwork, and manage their lifecycle (for resources that the authenticated user is associated with)

### Milestones

All milestones are returned on projects that the authenticated user is associated with. You can use the provided filters to return only those milestones that are incomplete, completed, upcoming or late.

### Risks

Manage project risks.

### Tasks / Tasklist

Manage projects tasks / taskslist.

Refer to the Teamwork connector [API specification](openapi.json) for details.

## Example Usage

< sdk example? >

## Authentication

To use the Teamwork Connector, you need your Teamwork URL and API Key.

Your URL can be found in the settings of your Teamwork site. If you navigate to the settings tab on the top right corner of your site and click on general. Your main project settings will appear you will see your URL there under 'Site Address'. You use this when pinging the API.

Your API token can be found by logging into your Teamwork account, clicking your avatar in the top right and choosing Edit my details. On the API tab of the dialog click the "Show your token" at the bottom (under "API Authentication tokens").

