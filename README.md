# api_runner

## Reason for being
When testing web APIs at work I either have to duplicate a request in bruno or repeatedly edit the same request, I often know the shape of what is expected and I will explore argument validation which would not be appropriate for fully automated end to end tests but may not have been covered in unit tests.
## Goals

### Project

- [x] run tests to speed up testing
- [x] small portable executable so I don't need set versions of dependencies like python
- [x] tests can be defined in a human readable format, JSON
- [x] user can define where the test file is and what it is called
- [x] produces a test report 
- [x] tests can replace part of default such as a URI

### Personal
- [x] learn some golang
- [x] understand if golang is suitable for other projects