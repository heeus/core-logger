# logger

Simple go logger. Default output will be like this:

```
09/24 12:36:56.951: **** ERR: [logger.Test_BasicUsage:26]: My error arg1 arg2
09/24 12:36:56.969: !!! WARN: [logger.Test_BasicUsage:27]: My warning
09/24 12:36:56.969: === INFO: [logger.Test_BasicUsage:28]: My info
09/24 12:36:56.969: --- VERB: [logger.Test_BasicUsage:39]: Now you should see my verbose
09/24 12:36:56.969: !!! WARN: [logger.Test_BasicUsage:45]: You should see my warning
09/24 12:36:56.969: !!! WARN: [logger.Test_BasicUsage:46]: You should see my info
09/24 12:36:56.969: **** ERR: [logger.(*mystruct).logMe:59]: OOPS
```

See [impl_test.Test_BasicUsage](impl_test.go#19) for examples
