# Agenda
Group members：

**15331151  李佳** 

**15331150  李辉旭** 

**15331145  李果** 

[![Build Status](https://travis-ci.org/Tomy-Lee/Agenda-WebService.svg?branch=master)](https://travis-ci.org/Tomy-Lee/Agenda-WebService)

![这里写图片描述](http://img.blog.csdn.net/20171211170850667?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzM0NTQxMTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171211170857432?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzM0NTQxMTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
### Installation and operation

```
docker pull Tomy-Lee/Agenda-WebService
docker run -dit --name agenda-sevice -v $PATH_TO_SERVER_DATA:/data -p 8080:8080 Tomy-Lee/Agenda-WebService service
docker run --rm --network host -v $PATH_TO_CLI_DATA:/data Tomy-Lee/Agenda-WebService cli help
```
- By setting the environment variable `  MOCK ` it accesses MOCK server.

- By using [``Travis CI``](https://travis-ci.org/Tomy-Lee/Agenda-WebService),it executes the go test command run writing good tests file for continuous integration.


### Client

#### Presentation Layer ``cmd``

- Responsible for accepting user input and handing over the business logic services provided by the business logic layer to get the results and show them to the user; Using ` fmt ` package to screen printing information.

#### Business Logic Layer ``service``

- Take care of simple form validation and call the interface provided by the entity layer to send the request to get the required data.

#### Entity layer ``entity``

- Responsible for interacting with the server, maintaining the data structure required for the interaction, and returning the data needed at the upper level.

### Service

service - dao - orm (entity) 

### Entity loading

- Entities use ` JSON ` format stored; Data loading and stored by the special ` storage ` structure is complete, be responsible for data manipulation ` Model using ` structure; Each ` Model ` should be only one instance, so consider using the singleton pattern. This is implemented using package global variables; Each ` Model ` load of IO operation, large amount of data is relatively time-consuming. And considering various ` Model ` the independence of the load, so we through ` goroutine ` parallel loading.

### log

- Each layer generates and uses different log instances through the factory mode, which can easily set the properties of the log required, such as the output destination, output format, etc. [after ` logrus `] (https://github.com/Sirupsen/logrus) package prefix is convenient to add; Record user operations and outputs in key areas of presentation, business logic, and physical layers.

### login control

- Using ` session ` said session, control the login state. Persisting in `  curUser. Json ` file；The validity of the session is controlled by the server. The client will put `openid ` on ` cookie` sent to the server to maintain a conversation.

### Exception handling

- For logic errors due to incorrect user input, use ` go ` language standard return error model, and prompt the user; For other force majeure cause abnormalities, such as file is written to failure, etc., using ` Panic ` function generating function stack and throw, in each ` goroutine ` and record to the outermost recovery ` log `, provide strong clues for debugging.
