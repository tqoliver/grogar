# grogar

## Deploy to OpenShift with the following:

### Create a new project
oc new-project grogar

### Create a new image stream
oc create is grogar-artifact

### Stage 1 Build: Create a build configuration for the builder

oc new-build https://github.com/cricci82/grogar --name=builder --to=grogar-artifact:latest

### Stage 2 Build: Create a build configuration for the runtime

oc new-build --name=runtime \\  
   --docker-image=scratch \\  
   --source-image=grogar-artifact \\  
   --source-image-path=/grogar:. \\  
   --dockerfile=$'FROM scratch\nCOPY /grogar /\nEXPOSE 8080\nUSER 1001\nENTRYPOINT ["/grogar"]'
   
