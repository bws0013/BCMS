This is a node that runs on a host machine
It monitors for changes in files that it is supposed to monitor
It also keeps a persistent list of changes that needs to be made

A rest api shall be used to pass information between services

Change process
  - Recieve change to be added
  - Add change to change list
  - Every N seconds close the list and announce that changes need to wait
  - For each entry in the list
    - Create directories if the option has been set
    - Retrieve intended file by communicating with one of the master nodes slaves
    - Delete/Change existing file
    - Store the md5 in persistent and dynamic storage
