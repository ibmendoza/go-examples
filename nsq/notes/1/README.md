```bash
0) ./nsqlookupd & ./nsqd --lookupd-tcp-address=127.0.0.1:4160 & ./nsqadmin --lookupd-http-address=127.0.0.1:4161
```

```bash
1) ./pub -msg=20

2015/12/01 07:55:22 INF    1 (192.168.56.101:4150) connecting to nsqd                                                                 
2015/12/01 07:55:22 Time took 43.259256ms                                                                                             
2015/12/01 07:55:22 INF    1 stopping                                                                                                 
2015/12/01 07:55:22 ERR    1 (192.168.56.101:4150) IO error - EOF                                                                     
2015/12/01 07:55:22 INF    1 (192.168.56.101:4150) beginning close                                                                    
2015/12/01 07:55:22 INF    1 (192.168.56.101:4150) readLoop exiting                                                                   
2015/12/01 07:55:22 INF    1 exiting router                                                                                           
2015/12/01 07:55:22 INF    1 (192.168.56.101:4150) breaking out of writeLoop                                                          
2015/12/01 07:55:22 INF    1 (192.168.56.101:4150) writeLoop exiting                                                                  
2015/12/01 07:55:22 INF    1 (192.168.56.101:4150) finished draining, cleanup exiting                                                 
2015/12/01 07:55:22 INF    1 (192.168.56.101:4150) clean close complete   
```

see output1.png

```bash
2) ./goq-witherror -msg=20

root@core /home/nsqapps# ./goq-witherror -msg=20                                                                                      
2015/12/01 08:43:13 INF    1 [test/ch] querying nsqlookupd http://:4161/lookup?topic=test                                             
2015/12/01 08:43:13 INF    1 [test/ch] (core:4150) connecting to nsqd                                                                 
2015/12/01 08:43:14 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce398c3000                                         
2015/12/01 08:43:14 WRN    1 [test/ch] backing off for 2.0000 seconds (backoff level 1), setting all to RDY 0                         
2015/12/01 08:43:14 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce398c3002                                         
2015/12/01 08:43:14 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce39cc3000                                         
2015/12/01 08:43:14 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce39cc3003                                         
2015/12/01 08:43:14 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce39cc3004                                         
2015/12/01 08:43:14 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce39cc3005                                         
2015/12/01 08:43:14 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce3a0c3003                                         
2015/12/01 08:43:14 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce3a0c3004                                         
2015/12/01 08:43:14 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce3a0c3005                                         
2015/12/01 08:43:14 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce3a0c3006                                         
2015/12/01 08:43:16 WRN    1 [test/ch] (core:4150) backoff timeout expired, sending RDY 1                                             
2015/12/01 08:43:29 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce398c3000                                         
2015/12/01 08:43:29 WRN    1 [test/ch] backing off for 4.0000 seconds (backoff level 2), setting all to RDY 0                         
2015/12/01 08:43:33 WRN    1 [test/ch] (core:4150) backoff timeout expired, sending RDY 1                                             
2015/12/01 08:43:33 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce398c3002                                         
2015/12/01 08:43:33 WRN    1 [test/ch] backing off for 8.0000 seconds (backoff level 3), setting all to RDY 0                         
2015/12/01 08:43:41 WRN    1 [test/ch] (core:4150) backoff timeout expired, sending RDY 1                                             
2015/12/01 08:43:41 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce39cc3000                                         
2015/12/01 08:43:41 WRN    1 [test/ch] backing off for 16.0000 seconds (backoff level 4), setting all to RDY 0                        
2015/12/01 08:43:57 WRN    1 [test/ch] (core:4150) backoff timeout expired, sending RDY 1                                             
2015/12/01 08:43:57 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce39cc3003                                         
2015/12/01 08:43:57 WRN    1 [test/ch] backing off for 32.0000 seconds (backoff level 5), setting all to RDY 0                        
2015/12/01 08:44:28 INF    1 [test/ch] querying nsqlookupd http://:4161/lookup?topic=test                                             
2015/12/01 08:44:29 WRN    1 [test/ch] (core:4150) backoff timeout expired, sending RDY 1                                             
2015/12/01 08:44:29 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce39cc3004                                         
2015/12/01 08:44:29 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce39cc3005                                         
2015/12/01 08:44:29 WRN    1 [test/ch] backing off for 64.0000 seconds (backoff level 6), setting all to RDY 0                        
2015/12/01 08:45:28 INF    1 [test/ch] querying nsqlookupd http://:4161/lookup?topic=test                                             
2015/12/01 08:45:33 WRN    1 [test/ch] (core:4150) backoff timeout expired, sending RDY 1                                             
2015/12/01 08:45:33 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce3a0c3003                                         
2015/12/01 08:45:33 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce3a0c3004                                         
2015/12/01 08:45:33 WRN    1 [test/ch] backing off for 64.0000 seconds (backoff level 6), setting all to RDY 0                        
2015/12/01 08:46:28 INF    1 [test/ch] querying nsqlookupd http://:4161/lookup?topic=test                                             
2015/12/01 08:46:37 WRN    1 [test/ch] (core:4150) backoff timeout expired, sending RDY 1                                             
2015/12/01 08:46:37 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce3a0c3005                                         
2015/12/01 08:46:37 WRN    1 [test/ch] backing off for 64.0000 seconds (backoff level 6), setting all to RDY 0                        
2015/12/01 08:47:28 INF    1 [test/ch] querying nsqlookupd http://:4161/lookup?topic=test                                             
2015/12/01 08:47:41 WRN    1 [test/ch] (core:4150) backoff timeout expired, sending RDY 1                                             
2015/12/01 08:47:41 ERR    1 [test/ch] Handler returned error (even) for msg 0951f3ce3a0c3006                                         
2015/12/01 08:47:41 WRN    1 [test/ch] backing off for 64.0000 seconds (backoff level 6), setting all to RDY 0        
```

see output2-5       

```
3) root@core /home/nsqapps# ./goq-noerror                                                                                        
2015/12/01 08:50:50 INF    1 [test/ch] querying nsqlookupd http://:4161/lookup?topic=test                                             
2015/12/01 08:50:50 INF    1 [test/ch] (core:4150) connecting to nsqd      
```

see output 6 to see all 20 pending messages in step2 have been all processed                                                                                                                                              
