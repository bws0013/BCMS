I had an idea for a basic config management system
Vainly I decided on calling is Ben's Config Management System

The central idea is something scalable that is inspired by some of the work on Rundeck

There are 2 libraries, one is your library of configs, the other is your nodes

You select a node and add a config to it, using key/value variable overrides if necessary

You also provide a directory location and if you want that directory created if it doesn't exist

The config file is added at that location

An agent on the machine stores the hashes of the config files, checking every N seconds for changed hashes. If they are different a call is made to one of the nodes in charge either to raise an alert or to retrieve a copy of the original to override the changes.

The priorities are scalability
  - Multiple master nodes will be available for calls to retrieve configs, slave nodes will actually send the configs
  - Databases should be mirrored HDFS style

Node groups and tags
  - This is inspired by Rundeck nodes and tags
  - There shall be system/node/project level variables
  - Key -> Value
