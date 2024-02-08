# aperture

## Iteration 0

Close the whole flow end to end - installable - skeleton stub

```shell
aperture init
```
- [x] Initialize .net app (default) 
  - [x] Run .net cli analyzer use docker in order to avoid multiple arch
  - [x] Run it

```shell
aperture run
```
- [ ] Build docker-compose configuration 
- [ ] Run all projections with .net runner (prebuild Docker image) keep it simple for now
- [ ] Aperture agent should be sending mock events via grpc
- [ ] Daemon collects simple stats somehow and it is mirrored in UI
- [ ] Aperture.sln pipeline to deploy core nuget and cli docker image

> At this point, we should be able to have a skeleton in place with well defined contracts and interfaces 
between different components and streamlined development process
 
## Iteration 1

Minimum usable thing - for evaluation purposes (eg. demoing and even standalone use by ppl)
 
--- 

Then build on this with additional provisioning and configuration (eg. db etc...), starting
stopping projections, restart etc ...