- http://jtfmumm.com/blog/2016/03/06/safely-sharing-data-pony-reference-capabilities/

**Laws of shared references**

1) Write Law: Write only when you know no one else can read or write.

2) Read Law: Read only when you know no one else can write.

The Write Law exists because, first, simultaneous writes to the same data structure can lead to unexpected write results and, second, writing while someone else is reading is likely to produce some weird read results. Whether by locks or reference capabilities, the Write Law must be enforced.
