# logger

Simple go logger with logging level. Default output will be like this:

```
09/28 17:40:57.064: **** ERR: [logger.Test_BasicUsage:26]: Hello world arg1 arg2
09/28 17:40:57.082: !!! WARN: [logger.Test_BasicUsage:27]: My warning
09/28 17:40:57.082: === INFO: [logger.Test_BasicUsage:28]: My info
09/28 17:40:57.082: --- DEBU: [logger.Test_BasicUsage:39]: Now you should see my Debug
09/28 17:40:57.082: !!! WARN: [logger.Test_BasicUsage:45]: You should see my warning
09/28 17:40:57.082: !!! WARN: [logger.Test_BasicUsage:46]: You should see my info
09/28 17:40:57.082: **** ERR: [logger.(*mystruct).logMe:59]: OOPS
```

See [impl_test.Test_BasicUsage](impl_test.go#L19) for examples

# Links

- [Why does the TRACE level exist, and when should I use it rather than DEBUG?](https://softwareengineering.stackexchange.com/questions/279690/why-does-the-trace-level-exist-and-when-should-i-use-it-rather-than-debug)
  - [Good answer](https://softwareengineering.stackexchange.com/questions/279690/why-does-the-trace-level-exist-and-when-should-i-use-it-rather-than-debug/360810#360810)