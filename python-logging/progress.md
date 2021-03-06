## [Logging HOWTO](https://docs.python.org/2/howto/logging.html)

- [x] Basic Logging Tutorial
    - [x] When to use logging
    - [x] A simple example
    - [x] Logging to a file
    - [x] Logging from multiple modules
    - [x] Logging variable data
    - [x] Changing the format of displayed messages
    - [x] Displaying the date/time in messages
    - [x] Next Steps
- [x] Advanced Logging Tutorial
    - [x] Logging Flow
    - [x] Loggers
    - [x] Handlers
    - [x] Formatters
    - [x] Configuring Logging
    - [x] What happens if no configuration is provided
    - [x] Configuring Logging for a Library
- [x] Logging Levels
    - [x] Custom Levels
- [x] Useful Handlers
- [x] Exceptions raised during logging
- [x] Using arbitrary objects as messages
- [x] Optimization

## [Logging Cookbook](https://docs.python.org/2/howto/logging-cookbook.html)

- [x] Using logging in multiple modules
- [x] Multiple handlers and formatters
- [x] Logging to multiple destinations
- [x] Configuration server example
- [x] Sending and receiving logging events across a network
- [x] Adding contextual information to your logging output
    - [x] Using LoggerAdapters to impart contextual information
        - [x] Using objects other than dicts to pass contextual information
    - [x] Using Filters to impart contextual information
- [x] Logging to a single file from multiple processes
- [x] Using file rotation
- [x] An example dictionary-based configuration
- [x] Inserting a BOM into messages sent to a SysLogHandler
- [x] Implementing structured logging
- [x] Customizing handlers with dictConfig()
- [x] Configuring filters with dictConfig()

## [logging — Logging facility for Python](https://docs.python.org/2/library/logging.html)

- [ ] Logger Objects
- [ ] Logging Levels
- [ ] Handler Objects
- [ ] Formatter Objects
- [ ] Filter Objects
- [ ] LogRecord Objects
- [ ] LogRecord attributes
- [ ] LoggerAdapter Objects
- [ ] Thread Safety
- [ ] Module-Level Functions
- [x] Integration with the warnings module

## [logging.config — Logging configuration](https://docs.python.org/2/library/logging.config.html)

- [ ] Configuration functions
- [ ] Configuration dictionary schema
    - [ ] Dictionary Schema Details
    - [ ] Incremental Configuration
    - [ ] Object connections
    - [ ] User-defined objects
    - [ ] Access to external objects
    - [ ] Access to internal objects
    - [ ] Import resolution and custom importers
- [ ] Configuration file format

## [logging.handlers — Logging handlers](https://docs.python.org/2/library/logging.handlers.html)

- [ ] StreamHandler
- [ ] FileHandler
- [ ] NullHandler
- [ ] WatchedFileHandler
- [ ] RotatingFileHandler
- [ ] TimedRotatingFileHandler
- [ ] SocketHandler
- [ ] DatagramHandler
- [ ] SysLogHandler
- [ ] NTEventLogHandler
- [ ] SMTPHandler
- [ ] MemoryHandler
- [ ] HTTPHandler

## [warnings — Warning control](https://docs.python.org/2/library/warnings.html)

- [x] Warning Categories
- [ ] The Warnings Filter
    - [ ] Default Warning Filters
- [ ] Temporarily Suppressing Warnings
- [ ] Testing Warnings
- [ ] Updating Code For New Versions of Python
- [ ] Available Functions
- [ ] Available Context Managers

## Other Reading

- [x] https://docs.python.org/2/tutorial/stdlib2.html#logging
- [ ] https://docs.python.org/2/library/multiprocessing.html#logging
- [ ] https://docs.python.org/2/whatsnew/2.3.html#pep-282-the-logging-package
- [ ] http://legacy.python.org/dev/peps/pep-0282/
- [ ] https://docs.python.org/2/whatsnew/2.7.html#pep-391-dictionary-based-configuration-for-logging
- [ ] http://legacy.python.org/dev/peps/pep-0391/
- [x] https://docs.python.org/2/whatsnew/2.1.html#pep-230-warning-framework
- [ ] http://legacy.python.org/dev/peps/pep-0230/
