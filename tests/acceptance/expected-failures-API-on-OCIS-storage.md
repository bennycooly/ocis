## Scenarios from ownCloud10 core API tests that are expected to fail with OCIS storage while running with the Graph API

The expected failures in this file are from features in the owncloud/ocis repo.

### File

Basic file management like up and download, move, copy, properties, trash, versions and chunking.

#### [Getting information about a folder overwritten by a file gives 500 error instead of 404](https://github.com/owncloud/ocis/issues/1239)

- [coreApiWebdavProperties1/copyFile.feature:274](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/copyFile.feature#L274)
- [coreApiWebdavProperties1/copyFile.feature:275](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/copyFile.feature#L275)
- [coreApiWebdavProperties1/copyFile.feature:292](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/copyFile.feature#L292)
- [coreApiWebdavProperties1/copyFile.feature:293](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/copyFile.feature#L293)

#### [Custom dav properties with namespaces are rendered incorrectly](https://github.com/owncloud/ocis/issues/2140)

_ocdav: double-check the webdav property parsing when custom namespaces are used_

- [coreApiWebdavProperties1/setFileProperties.feature:37](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/setFileProperties.feature#L37)
- [coreApiWebdavProperties1/setFileProperties.feature:38](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/setFileProperties.feature#L38)
- [coreApiWebdavProperties1/setFileProperties.feature:43](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/setFileProperties.feature#L43)
- [coreApiWebdavProperties1/setFileProperties.feature:78](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/setFileProperties.feature#L78)
- [coreApiWebdavProperties1/setFileProperties.feature:79](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/setFileProperties.feature#L79)
- [coreApiWebdavProperties1/setFileProperties.feature:84](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/setFileProperties.feature#L84)

#### [Cannot set custom webDav properties](https://github.com/owncloud/product/issues/264)

- [coreApiWebdavProperties2/getFileProperties.feature:360](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties2/getFileProperties.feature#L360)
- [coreApiWebdavProperties2/getFileProperties.feature:365](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties2/getFileProperties.feature#L365)
- [coreApiWebdavProperties2/getFileProperties.feature:370](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties2/getFileProperties.feature#L370)
- [coreApiWebdavProperties2/getFileProperties.feature:400](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties2/getFileProperties.feature#L400)
- [coreApiWebdavProperties2/getFileProperties.feature:405](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties2/getFileProperties.feature#L405)
- [coreApiWebdavProperties2/getFileProperties.feature:410](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties2/getFileProperties.feature#L410)

#### [file versions do not report the version author](https://github.com/owncloud/ocis/issues/2914)

- [coreApiVersions/fileVersionAuthor.feature:14](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiVersions/fileVersionAuthor.feature#L14)
- [coreApiVersions/fileVersionAuthor.feature:37](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiVersions/fileVersionAuthor.feature#L37)
- [coreApiVersions/fileVersionAuthor.feature:58](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiVersions/fileVersionAuthor.feature#L58)
- [coreApiVersions/fileVersionAuthor.feature:78](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiVersions/fileVersionAuthor.feature#L78)
- [coreApiVersions/fileVersionAuthor.feature:104](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiVersions/fileVersionAuthor.feature#L104)
- [coreApiVersions/fileVersionAuthor.feature:129](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiVersions/fileVersionAuthor.feature#L129)
- [coreApiVersions/fileVersionAuthor.feature:154](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiVersions/fileVersionAuthor.feature#L154)
- [coreApiVersions/fileVersionAuthor.feature:180](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiVersions/fileVersionAuthor.feature#L180)
- [coreApiVersions/fileVersionAuthor.feature:223](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiVersions/fileVersionAuthor.feature#L223)

### Sync

Synchronization features like etag propagation, setting mtime and locking files

#### [Uploading an old method chunked file with checksum should fail using new DAV path](https://github.com/owncloud/ocis/issues/2323)

- [coreApiMain/checksums.feature:369](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiMain/checksums.feature#L369)
- [coreApiMain/checksums.feature:374](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiMain/checksums.feature#L374)

#### [Webdav LOCK operations](https://github.com/owncloud/ocis/issues/1284)

- [coreApiWebdavLocks/exclusiveLocks.feature:124](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L124)
- [coreApiWebdavLocks/exclusiveLocks.feature:125](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L125)
- [coreApiWebdavLocks/exclusiveLocks.feature:126](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L126)
- [coreApiWebdavLocks/exclusiveLocks.feature:127](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L127)
- [coreApiWebdavLocks/exclusiveLocks.feature:132](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L132)
- [coreApiWebdavLocks/exclusiveLocks.feature:133](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L133)
- [coreApiWebdavLocks/exclusiveLocks.feature:151](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L151)
- [coreApiWebdavLocks/exclusiveLocks.feature:152](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L152)
- [coreApiWebdavLocks/exclusiveLocks.feature:153](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L153)
- [coreApiWebdavLocks/exclusiveLocks.feature:154](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L154)
- [coreApiWebdavLocks/exclusiveLocks.feature:159](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L159)
- [coreApiWebdavLocks/exclusiveLocks.feature:160](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L160)
- [coreApiWebdavLocks/exclusiveLocks.feature:178](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L178)
- [coreApiWebdavLocks/exclusiveLocks.feature:179](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L179)
- [coreApiWebdavLocks/exclusiveLocks.feature:180](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L180)
- [coreApiWebdavLocks/exclusiveLocks.feature:181](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L181)
- [coreApiWebdavLocks/exclusiveLocks.feature:186](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L186)
- [coreApiWebdavLocks/exclusiveLocks.feature:187](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/exclusiveLocks.feature#L187)
- [coreApiWebdavLocks/requestsWithToken.feature:126](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/requestsWithToken.feature#L126)
- [coreApiWebdavLocks/requestsWithToken.feature:127](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/requestsWithToken.feature#L127)
- [coreApiWebdavLocks/requestsWithToken.feature:132](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks/requestsWithToken.feature#L132)
- [coreApiWebdavLocks3/independentLocks.feature:65](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L65)
- [coreApiWebdavLocks3/independentLocks.feature:66](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L66)
- [coreApiWebdavLocks3/independentLocks.feature:67](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L67)
- [coreApiWebdavLocks3/independentLocks.feature:68](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L68)
- [coreApiWebdavLocks3/independentLocks.feature:73](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L73)
- [coreApiWebdavLocks3/independentLocks.feature:74](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L74)
- [coreApiWebdavLocks3/independentLocks.feature:93](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L93)
- [coreApiWebdavLocks3/independentLocks.feature:94](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L94)
- [coreApiWebdavLocks3/independentLocks.feature:95](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L95)
- [coreApiWebdavLocks3/independentLocks.feature:96](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L96)
- [coreApiWebdavLocks3/independentLocks.feature:97](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L97)
- [coreApiWebdavLocks3/independentLocks.feature:98](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L98)
- [coreApiWebdavLocks3/independentLocks.feature:99](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L99)
- [coreApiWebdavLocks3/independentLocks.feature:100](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L100)
- [coreApiWebdavLocks3/independentLocks.feature:105](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L105)
- [coreApiWebdavLocks3/independentLocks.feature:106](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L106)
- [coreApiWebdavLocks3/independentLocks.feature:107](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L107)
- [coreApiWebdavLocks3/independentLocks.feature:108](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocks.feature#L108)
- [coreApiWebdavLocks3/independentLocksShareToShares.feature:75](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocksShareToShares.feature#L75)
- [coreApiWebdavLocks3/independentLocksShareToShares.feature:76](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocksShareToShares.feature#L76)
- [coreApiWebdavLocks3/independentLocksShareToShares.feature:77](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocksShareToShares.feature#L77)
- [coreApiWebdavLocks3/independentLocksShareToShares.feature:78](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocksShareToShares.feature#L78)
- [coreApiWebdavLocks3/independentLocksShareToShares.feature:83](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocksShareToShares.feature#L83)
- [coreApiWebdavLocks3/independentLocksShareToShares.feature:84](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocksShareToShares.feature#L84)
- [coreApiWebdavLocks3/independentLocksShareToShares.feature:104](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocksShareToShares.feature#L104)
- [coreApiWebdavLocks3/independentLocksShareToShares.feature:105](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocksShareToShares.feature#L105)
- [coreApiWebdavLocks3/independentLocksShareToShares.feature:106](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocksShareToShares.feature#L106)
- [coreApiWebdavLocks3/independentLocksShareToShares.feature:107](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocksShareToShares.feature#L107)
- [coreApiWebdavLocks3/independentLocksShareToShares.feature:112](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocksShareToShares.feature#L112)
- [coreApiWebdavLocks3/independentLocksShareToShares.feature:113](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocks3/independentLocksShareToShares.feature#L113)
- [coreApiWebdavLocksUnlock/unlock.feature:40](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlock.feature#L40)
- [coreApiWebdavLocksUnlock/unlock.feature:41](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlock.feature#L41)
- [coreApiWebdavLocksUnlock/unlock.feature:46](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlock.feature#L46)
- [coreApiWebdavLocksUnlock/unlock.feature:79](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlock.feature#L79)
- [coreApiWebdavLocksUnlock/unlock.feature:80](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlock.feature#L80)
- [coreApiWebdavLocksUnlock/unlock.feature:129](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlock.feature#L129)
- [coreApiWebdavLocksUnlock/unlock.feature:130](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlock.feature#L130)
- [coreApiWebdavLocksUnlock/unlock.feature:131](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlock.feature#L131)
- [coreApiWebdavLocksUnlock/unlock.feature:132](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlock.feature#L132)
- [coreApiWebdavLocksUnlock/unlock.feature:137](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlock.feature#L137)
- [coreApiWebdavLocksUnlock/unlock.feature:138](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlock.feature#L138)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:27](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L27)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:28](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L28)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:29](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L29)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:30](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L30)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:35](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L35)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:36](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L36)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:51](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L51)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:52](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L52)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:53](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L53)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:54](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L54)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:59](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L59)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:60](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L60)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:99](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L99)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:100](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L100)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:101](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L101)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:102](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L102)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:107](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L107)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:108](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L108)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:123](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L123)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:124](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L124)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:125](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L125)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:126](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L126)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:131](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L131)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:132](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L132)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:147](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L147)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:148](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L148)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:149](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L149)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:150](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L150)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:155](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L155)
- [coreApiWebdavLocksUnlock/unlockSharingToShares.feature:156](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavLocksUnlock/unlockSharingToShares.feature#L156)

### Share

File and sync features in a shared scenario

#### [ocs sharing api always returns an empty exact list while searching for a sharee](https://github.com/owncloud/ocis/issues/4265)

- [coreApiSharees/sharees.feature:350](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharees/sharees.feature#L350)
- [coreApiSharees/sharees.feature:351](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharees/sharees.feature#L351)
- [coreApiSharees/sharees.feature:370](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharees/sharees.feature#L370)
- [coreApiSharees/sharees.feature:371](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharees/sharees.feature#L371)
- [coreApiSharees/sharees.feature:390](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharees/sharees.feature#L390)
- [coreApiSharees/sharees.feature:391](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharees/sharees.feature#L391)
- [coreApiSharees/sharees.feature:410](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharees/sharees.feature#L410)
- [coreApiSharees/sharees.feature:411](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharees/sharees.feature#L411)
- [coreApiSharees/sharees.feature:430](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharees/sharees.feature#L430)
- [coreApiSharees/sharees.feature:431](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharees/sharees.feature#L431)


#### [accepting matching name shared resources from different users/groups sets no serial identifiers on the resource name for the receiver](https://github.com/owncloud/ocis/issues/4289)

- [coreApiShareManagementToShares/acceptShares.feature:334](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/acceptShares.feature#L334)
- [coreApiShareManagementToShares/acceptShares.feature:364](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/acceptShares.feature#L364)
- [coreApiShareManagementToShares/acceptShares.feature:278](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/acceptShares.feature#L278)
- [coreApiShareManagementToShares/acceptShares.feature:543](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/acceptShares.feature#L543)
- [coreApiShareManagementToShares/acceptShares.feature:608](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/acceptShares.feature#L608)
- [coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature:141](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature#L141)
- [coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature:142](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature#L142)
- [coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature:181](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature#L181)
- [coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature:182](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature#L182)
- [coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature:45](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature#L45)
- [coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature:46](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature#L46)

#### [sharing the shares folder to users exits with different status code than in oc10 backend](https://github.com/owncloud/ocis/issues/2215)

- [coreApiShareManagementBasicToShares/createShareToSharesFolder.feature:747](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/createShareToSharesFolder.feature#L747)
- [coreApiShareManagementBasicToShares/createShareToSharesFolder.feature:748](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/createShareToSharesFolder.feature#L748)
- [coreApiShareManagementBasicToShares/createShareToSharesFolder.feature:766](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/createShareToSharesFolder.feature#L766)
- [coreApiShareManagementBasicToShares/createShareToSharesFolder.feature:767](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/createShareToSharesFolder.feature#L767)
- [coreApiShareManagementBasicToShares/createShareToSharesFolder.feature:782](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/createShareToSharesFolder.feature#L782)
- [coreApiShareManagementBasicToShares/createShareToSharesFolder.feature:783](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/createShareToSharesFolder.feature#L783)

#### [file_target of an auto-renamed file is not correct directly after sharing](https://github.com/owncloud/core/issues/32322)

- [coreApiShareManagementToShares/mergeShare.feature:121](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/mergeShare.feature#L121)

#### [File deletion using dav gives unique string in filename in the trashbin](https://github.com/owncloud/product/issues/178)

- [coreApiShareManagementBasicToShares/deleteShareFromShares.feature:63](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/deleteShareFromShares.feature#L63)
- [coreApiShareManagementBasicToShares/deleteShareFromShares.feature:77](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/deleteShareFromShares.feature#L77)

cannot share a folder with create permission

#### [Resource with share permission create is readable for sharee](https://github.com/owncloud/ocis/issues/4524)

- [coreApiShareManagementBasicToShares/deleteShareFromShares.feature:130](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/deleteShareFromShares.feature#L130)
- [coreApiShareManagementBasicToShares/deleteShareFromShares.feature:143](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/deleteShareFromShares.feature#L143)


#### [OCS error message for attempting to access share via share id as an unauthorized user is not informative](https://github.com/owncloud/ocis/issues/1233)

- [coreApiShareOperationsToShares1/gettingShares.feature:184](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares1/gettingShares.feature#L184)
- [coreApiShareOperationsToShares1/gettingShares.feature:185](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares1/gettingShares.feature#L185)

#### [Listing shares via ocs API does not show path for parent folders](https://github.com/owncloud/ocis/issues/1231)

- [coreApiShareOperationsToShares1/gettingShares.feature:221](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares1/gettingShares.feature#L221)
- [coreApiShareOperationsToShares1/gettingShares.feature:222](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares1/gettingShares.feature#L222)

#### [Public link enforce permissions](https://github.com/owncloud/ocis/issues/1269)

- [coreApiSharePublicLink1/accessToPublicLinkShare.feature:10](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink1/accessToPublicLinkShare.feature#L10)
- [coreApiSharePublicLink1/accessToPublicLinkShare.feature:21](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink1/accessToPublicLinkShare.feature#L21)
- [coreApiSharePublicLink1/accessToPublicLinkShare.feature:31](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink1/accessToPublicLinkShare.feature#L31)
- [coreApiSharePublicLink1/accessToPublicLinkShare.feature:47](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink1/accessToPublicLinkShare.feature#L47)
- [coreApiSharePublicLink1/createPublicLinkShare.feature:528](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink1/createPublicLinkShare.feature#L528)
- [coreApiSharePublicLink1/createPublicLinkShare.feature:549](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink1/createPublicLinkShare.feature#L549)

#### [download previews of other users file](https://github.com/owncloud/ocis/issues/2071)

- [coreApiWebdavPreviews/previews.feature:102](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavPreviews/previews.feature#L102)

#### [different error message detail for previews of folder](https://github.com/owncloud/ocis/issues/2064)

- [coreApiWebdavPreviews/previews.feature:111](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavPreviews/previews.feature#L111)

#### [Requesting a file preview when it is disabled by the administrator](https://github.com/owncloud/ocis/issues/192)

- [coreApiWebdavPreviews/previews.feature:126](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavPreviews/previews.feature#L126)

#### [Cannot set/unset maximum and minimum preview dimensions](https://github.com/owncloud/ocis/issues/2070)

- [coreApiWebdavPreviews/previews.feature:134](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavPreviews/previews.feature#L134)
- [coreApiWebdavPreviews/previews.feature:162](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavPreviews/previews.feature#L162)
- [coreApiWebdavPreviews/previews.feature:163](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavPreviews/previews.feature#L163)
- [coreApiWebdavPreviews/previews.feature:164](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavPreviews/previews.feature#L164)
- [coreApiWebdavPreviews/previews.feature:176](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavPreviews/previews.feature#L176)
- [coreApiWebdavPreviews/previews.feature:177](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavPreviews/previews.feature#L177)

#### [creating public links with permissions fails](https://github.com/owncloud/product/issues/252)

- [coreApiSharePublicLink1/changingPublicLinkShare.feature:30](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink1/changingPublicLinkShare.feature#L30)
- [coreApiSharePublicLink1/changingPublicLinkShare.feature:51](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink1/changingPublicLinkShare.feature#L51)
- [coreApiSharePublicLink1/changingPublicLinkShare.feature:90](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink1/changingPublicLinkShare.feature#L90)

#### [copying a folder within a public link folder to folder with same name as an already existing file overwrites the parent file](https://github.com/owncloud/ocis/issues/1232)

- [coreApiSharePublicLink2/copyFromPublicLink.feature:63](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink2/copyFromPublicLink.feature#L63)
- [coreApiSharePublicLink2/copyFromPublicLink.feature:89](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink2/copyFromPublicLink.feature#L89)
- [coreApiSharePublicLink2/copyFromPublicLink.feature:173](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink2/copyFromPublicLink.feature#L173)
- [coreApiSharePublicLink2/copyFromPublicLink.feature:174](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink2/copyFromPublicLink.feature#L174)
- [coreApiSharePublicLink2/copyFromPublicLink.feature:189](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink2/copyFromPublicLink.feature#L189)
- [coreApiSharePublicLink2/copyFromPublicLink.feature:190](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink2/copyFromPublicLink.feature#L190)
- [coreApiSharePublicLink3/updatePublicLinkShare.feature:45](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink3/updatePublicLinkShare.feature#L45)
- [coreApiSharePublicLink3/updatePublicLinkShare.feature:46](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink3/updatePublicLinkShare.feature#L46)

#### [Upload-only shares must not overwrite but create a separate file](https://github.com/owncloud/ocis/issues/1267)

- [coreApiSharePublicLink3/uploadToPublicLinkShare.feature:24](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink3/uploadToPublicLinkShare.feature#L24)
- [coreApiSharePublicLink3/uploadToPublicLinkShare.feature:273](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink3/uploadToPublicLinkShare.feature#L273)

#### [Set quota over settings](https://github.com/owncloud/ocis/issues/1290)

- [coreApiSharePublicLink3/uploadToPublicLinkShare.feature:160](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink3/uploadToPublicLinkShare.feature#L160)
- [coreApiSharePublicLink3/uploadToPublicLinkShare.feature:179](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink3/uploadToPublicLinkShare.feature#L179)


#### [deleting a file inside a received shared folder is moved to the trash-bin of the sharer not the receiver](https://github.com/owncloud/ocis/issues/1124)

- [coreApiTrashbin/trashbinSharingToShares.feature:29](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinSharingToShares.feature#L29)
- [coreApiTrashbin/trashbinSharingToShares.feature:46](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinSharingToShares.feature#L46)
- [coreApiTrashbin/trashbinSharingToShares.feature:51](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinSharingToShares.feature#L51)
- [coreApiTrashbin/trashbinSharingToShares.feature:73](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinSharingToShares.feature#L73)
- [coreApiTrashbin/trashbinSharingToShares.feature:78](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinSharingToShares.feature#L78)
- [coreApiTrashbin/trashbinSharingToShares.feature:100](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinSharingToShares.feature#L100)
- [coreApiTrashbin/trashbinSharingToShares.feature:105](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinSharingToShares.feature#L105)
- [coreApiTrashbin/trashbinSharingToShares.feature:128](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinSharingToShares.feature#L128)
- [coreApiTrashbin/trashbinSharingToShares.feature:133](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinSharingToShares.feature#L133)
- [coreApiTrashbin/trashbinSharingToShares.feature:156](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinSharingToShares.feature#L156)
- [coreApiTrashbin/trashbinSharingToShares.feature:161](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinSharingToShares.feature#L161)
- [coreApiTrashbin/trashbinSharingToShares.feature:184](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinSharingToShares.feature#L184)
- [coreApiTrashbin/trashbinSharingToShares.feature:189](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinSharingToShares.feature#L189)
- [coreApiTrashbin/trashbinSharingToShares.feature:212](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinSharingToShares.feature#L212)
- [coreApiTrashbin/trashbinSharingToShares.feature:236](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinSharingToShares.feature#L236)

#### [changing user quota gives ocs status 103 / Cannot set quota](https://github.com/owncloud/product/issues/247)

- [coreApiShareOperationsToShares2/uploadToShare.feature:210](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/uploadToShare.feature#L210)
- [coreApiShareOperationsToShares2/uploadToShare.feature:211](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/uploadToShare.feature#L211)

#### [not possible to move file into a received folder](https://github.com/owncloud/ocis/issues/764)

- [coreApiShareOperationsToShares1/changingFilesShare.feature:25](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares1/changingFilesShare.feature#L25)
- [coreApiShareOperationsToShares1/changingFilesShare.feature:26](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares1/changingFilesShare.feature#L26)
- [coreApiShareOperationsToShares1/changingFilesShare.feature:109](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares1/changingFilesShare.feature#L109)
- [coreApiShareOperationsToShares1/changingFilesShare.feature:110](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares1/changingFilesShare.feature#L110)
- [coreApiShareOperationsToShares1/changingFilesShare.feature:131](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares1/changingFilesShare.feature#L131)
- [coreApiShareOperationsToShares1/changingFilesShare.feature:132](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares1/changingFilesShare.feature#L132)
- [coreApiShareManagementBasicToShares/createShareToSharesFolder.feature:538](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/createShareToSharesFolder.feature#L538)
- [coreApiVersions/fileVersionsSharingToShares.feature:220](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiVersions/fileVersionsSharingToShares.feature#L220)
- [coreApiVersions/fileVersionsSharingToShares.feature:221](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiVersions/fileVersionsSharingToShares.feature#L221)
- [coreApiWebdavMove2/moveShareOnOcis.feature:30](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveShareOnOcis.feature#L30)
- [coreApiWebdavMove2/moveShareOnOcis.feature:32](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveShareOnOcis.feature#L32)
- [coreApiWebdavMove2/moveShareOnOcis.feature:98](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveShareOnOcis.feature#L98)
- [coreApiWebdavMove2/moveShareOnOcis.feature:100](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveShareOnOcis.feature#L100)
- [coreApiWebdavMove2/moveShareOnOcis.feature:169](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveShareOnOcis.feature#L169)
- [coreApiWebdavMove2/moveShareOnOcis.feature:170](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveShareOnOcis.feature#L170)

#### [Expiration date for shares is not implemented](https://github.com/owncloud/ocis/issues/1250)

#### Expiration date of user shares

- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:52](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L52)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:53](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L53)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:76](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L76)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:77](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L77)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:102](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L102)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:103](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L103)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:128](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L128)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:129](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L129)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:279](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L279)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:280](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L280)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:301](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L301)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:302](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L302)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:323](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L323)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:324](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L324)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:346](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L346)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:347](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L347)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:363](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L363)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:364](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L364)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:380](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L380)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:381](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L381)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:576](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L576)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:577](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L577)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:599](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L599)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:600](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L600)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:601](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L601)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:602](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L602)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:603](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L603)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:624](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L624)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:625](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L625)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:626](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L626)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:627](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L627)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:628](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L628)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:629](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L629)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:630](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L630)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:631](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L631)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:632](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L632)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:633](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L633)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:634](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L634)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:635](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L635)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:656](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L656)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:657](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L657)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:658](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L658)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:659](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L659)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:660](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L660)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:661](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L661)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:682](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L682)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:683](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L683)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:684](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L684)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:685](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L685)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:686](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L686)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:687](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L687)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:708](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L708)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:709](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L709)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:732](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L732)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:733](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L733)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:756](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L756)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:757](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L757)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:34](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L34)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:35](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L35)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:86](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L86)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:87](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L87)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:143](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L143)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:144](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L144)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:201](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L201)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:202](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L202)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:203](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L203)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:204](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L204)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:287](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L287)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:288](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L288)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:318](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L318)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:319](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L319)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:320](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L320)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:321](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L321)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:379](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L379)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:380](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L380)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:381](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L381)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:382](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L382)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:383](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L383)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:384](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L384)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:413](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L413)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:414](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L414)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:415](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L415)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:416](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L416)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:444](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L444)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:445](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L445)

#### Expiration date of group shares

- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:175](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L175)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:176](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L176)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:201](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L201)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:202](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L202)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:229](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L229)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:230](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L230)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:258](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L258)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:259](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L259)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:403](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L403)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:404](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L404)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:427](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L427)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:428](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L428)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:451](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L451)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:452](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L452)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:476](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L476)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:477](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L477)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:497](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L497)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:498](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L498)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:518](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L518)
- [coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature:519](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares1/createShareExpirationDate.feature#L519)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:60](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L60)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:61](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L61)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:116](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L116)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:117](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L117)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:172](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L172)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:173](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L173)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:232](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L232)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:233](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L233)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:234](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L234)
- [coreApiShareReshareToShares3/reShareWithExpiryDate.feature:235](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareReshareToShares3/reShareWithExpiryDate.feature#L235)

#### [Cannot move folder/file from one received share to another](https://github.com/owncloud/ocis/issues/2442)

- [coreApiShareUpdateToShares/updateShare.feature:235](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareUpdateToShares/updateShare.feature#L235)
- [coreApiShareUpdateToShares/updateShare.feature:194](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareUpdateToShares/updateShare.feature#L194)
- [coreApiShareManagementToShares/mergeShare.feature:141](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/mergeShare.feature#L141)

#### [Sharing folder and sub-folder with same user but different permission,the permission of sub-folder is not obeyed ](https://github.com/owncloud/ocis/issues/2440)

- [coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature:278](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature#L278)
- [coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature:313](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature#L313)
- [coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature:422](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature#L422)
- [coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature:457](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares2/createShareReceivedInMultipleWays.feature#L457)

#### [Empty OCS response for a share create request using a disabled user](https://github.com/owncloud/ocis/issues/2212)

- [coreApiShareCreateSpecialToShares2/createShareWithDisabledUser.feature:20](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares2/createShareWithDisabledUser.feature#L20)
- [coreApiShareCreateSpecialToShares2/createShareWithDisabledUser.feature:23](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareCreateSpecialToShares2/createShareWithDisabledUser.feature#L23)


#### [Edit user share response has a "name" field](https://github.com/owncloud/ocis/issues/1225)

- [coreApiShareUpdateToShares/updateShare.feature:281](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareUpdateToShares/updateShare.feature#L281)
- [coreApiShareUpdateToShares/updateShare.feature:282](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareUpdateToShares/updateShare.feature#L282)

#### [user can access version metadata of a received share before accepting it](https://github.com/owncloud/ocis/issues/760)

- [coreApiVersions/fileVersionsSharingToShares.feature:283](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiVersions/fileVersionsSharingToShares.feature#L283)

#### [Share lists deleted user as 'user'](https://github.com/owncloud/ocis/issues/903)

- [coreApiShareManagementBasicToShares/createShareToSharesFolder.feature:682](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/createShareToSharesFolder.feature#L682)
- [coreApiShareManagementBasicToShares/createShareToSharesFolder.feature:683](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/createShareToSharesFolder.feature#L683)

#### [deleting a share with wrong authentication returns OCS status 996 / HTTP 500](https://github.com/owncloud/ocis/issues/1229)

- [coreApiShareManagementBasicToShares/deleteShareFromShares.feature:226](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/deleteShareFromShares.feature#L226)
- [coreApiShareManagementBasicToShares/deleteShareFromShares.feature:227](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/deleteShareFromShares.feature#L227)

### User Management

User and group management features

#### [incorrect ocs(v2) status value when getting info of share that does not exist should be 404, gives 998](https://github.com/owncloud/product/issues/250)

_ocs: api compatibility, return correct status code_

- [coreApiShareOperationsToShares2/shareAccessByID.feature:48](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L48)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:49](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L49)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:50](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L50)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:51](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L51)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:52](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L52)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:53](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L53)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:54](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L54)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:55](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L55)

### Other

API, search, favorites, config, capabilities, not existing endpoints, CORS and others

#### [Ability to return error messages in Webdav response bodies](https://github.com/owncloud/ocis/issues/1293)

- [coreApiAuthOcs/ocsDELETEAuth.feature:8](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthOcs/ocsDELETEAuth.feature#L8)
- [coreApiAuthOcs/ocsGETAuth.feature:8](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthOcs/ocsGETAuth.feature#L8)
- [coreApiAuthOcs/ocsGETAuth.feature:42](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthOcs/ocsGETAuth.feature#L42)
- [coreApiAuthOcs/ocsGETAuth.feature:73](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthOcs/ocsGETAuth.feature#L73)
- [coreApiAuthOcs/ocsGETAuth.feature:104](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthOcs/ocsGETAuth.feature#L104)
- [coreApiAuthOcs/ocsGETAuth.feature:121](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthOcs/ocsGETAuth.feature#L121)
- [coreApiAuthOcs/ocsPOSTAuth.feature:8](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthOcs/ocsPOSTAuth.feature#L8)
- [coreApiAuthOcs/ocsPUTAuth.feature:8](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthOcs/ocsPUTAuth.feature#L8)

#### [sending MKCOL requests to another user's webDav endpoints as normal user gives 404 instead of 403 ](https://github.com/owncloud/ocis/issues/3872)

_ocdav: api compatibility, return correct status code_

- [coreApiAuthWebDav/webDavMKCOLAuth.feature:52](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavMKCOLAuth.feature#L52)
- [coreApiAuthWebDav/webDavMKCOLAuth.feature:63](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavMKCOLAuth.feature#L63)

#### [trying to lock file of another user gives http 200](https://github.com/owncloud/ocis/issues/2176)

- [coreApiAuthWebDav/webDavLOCKAuth.feature:56](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavLOCKAuth.feature#L56)
- [coreApiAuthWebDav/webDavLOCKAuth.feature:68](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavLOCKAuth.feature#L68)

#### [send (MOVE,COPY) requests to another user's webDav endpoints as normal user gives 400 instead of 403](https://github.com/owncloud/ocis/issues/3882)

_ocdav: api compatibility, return correct status code_

- [coreApiAuthWebDav/webDavMOVEAuth.feature:55](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavMOVEAuth.feature#L55)
- [coreApiAuthWebDav/webDavMOVEAuth.feature:64](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavMOVEAuth.feature#L64)
- [coreApiAuthWebDav/webDavCOPYAuth.feature:59](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavCOPYAuth.feature#L59)
- [coreApiAuthWebDav/webDavCOPYAuth.feature:68](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavCOPYAuth.feature#L68)

#### [send POST requests to another user's webDav endpoints as normal user](https://github.com/owncloud/ocis/issues/1287)

_ocdav: api compatibility, return correct status code_

- [coreApiAuthWebDav/webDavPOSTAuth.feature:56](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavPOSTAuth.feature#L56)
- [coreApiAuthWebDav/webDavPOSTAuth.feature:65](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavPOSTAuth.feature#L65)

#### Another users space literally does not exist because it is not listed as a space for him, 404 seems correct, expects 403

- [coreApiAuthWebDav/webDavPUTAuth.feature:56](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavPUTAuth.feature#L56)
- [coreApiAuthWebDav/webDavPUTAuth.feature:68](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavPUTAuth.feature#L68)

#### [Using double slash in URL to access a folder gives 501 and other status codes](https://github.com/owncloud/ocis/issues/1667)

- [coreApiAuthWebDav/webDavSpecialURLs.feature:13](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavSpecialURLs.feature#L13)
- [coreApiAuthWebDav/webDavSpecialURLs.feature:24](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavSpecialURLs.feature#L24)
- [coreApiAuthWebDav/webDavSpecialURLs.feature:55](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavSpecialURLs.feature#L55)
- [coreApiAuthWebDav/webDavSpecialURLs.feature:66](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavSpecialURLs.feature#L66)
- [coreApiAuthWebDav/webDavSpecialURLs.feature:76](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavSpecialURLs.feature#L76)
- [coreApiAuthWebDav/webDavSpecialURLs.feature:88](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavSpecialURLs.feature#L88)
- [coreApiAuthWebDav/webDavSpecialURLs.feature:100](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavSpecialURLs.feature#L100)
- [coreApiAuthWebDav/webDavSpecialURLs.feature:111](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavSpecialURLs.feature#L111)
- [coreApiAuthWebDav/webDavSpecialURLs.feature:121](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavSpecialURLs.feature#L121)
- [coreApiAuthWebDav/webDavSpecialURLs.feature:132](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavSpecialURLs.feature#L132)
- [coreApiAuthWebDav/webDavSpecialURLs.feature:142](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavSpecialURLs.feature#L142)
- [coreApiAuthWebDav/webDavSpecialURLs.feature:153](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavSpecialURLs.feature#L153)
- [coreApiAuthWebDav/webDavSpecialURLs.feature:163](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavSpecialURLs.feature#L163)
- [coreApiAuthWebDav/webDavSpecialURLs.feature:174](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavSpecialURLs.feature#L174)
- [coreApiAuthWebDav/webDavSpecialURLs.feature:184](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavSpecialURLs.feature#L184)
- [coreApiAuthWebDav/webDavSpecialURLs.feature:195](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavSpecialURLs.feature#L195)

#### [Difference in response content of status.php and default capabilities](https://github.com/owncloud/ocis/issues/1286)

- [coreApiCapabilities/capabilitiesWithNormalUser.feature:11](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiCapabilities/capabilitiesWithNormalUser.feature#L11)

#### [[old/new/spaces] In ocis and oc10, REPORT request response differently](https://github.com/owncloud/ocis/issues/4712)

- [coreApiWebdavOperations/search.feature:208](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/search.feature#L208)
- [coreApiWebdavOperations/search.feature:209](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/search.feature#L209)
- [coreApiWebdavOperations/search.feature:214](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/search.feature#L214)
- [coreApiWebdavOperations/search.feature:240](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/search.feature#L240)
- [coreApiWebdavOperations/search.feature:241](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/search.feature#L241)
- [coreApiWebdavOperations/search.feature:246](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/search.feature#L246)

#### [Support for favorites](https://github.com/owncloud/ocis/issues/1228)

- [coreApiFavorites/favorites.feature:115](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favorites.feature#L115)
- [coreApiFavorites/favorites.feature:116](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favorites.feature#L116)
- [coreApiFavorites/favorites.feature:142](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favorites.feature#L142)
- [coreApiFavorites/favorites.feature:143](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favorites.feature#L143)
- [coreApiFavorites/favorites.feature:264](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favorites.feature#L264)
- [coreApiFavorites/favorites.feature:265](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favorites.feature#L265)

And other missing implementation of favorites

- [coreApiFavorites/favorites.feature:183](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favorites.feature#L183)
- [coreApiFavorites/favorites.feature:184](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favorites.feature#L184)
- [coreApiFavorites/favorites.feature:189](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favorites.feature#L189)
- [coreApiFavorites/favorites.feature:216](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favorites.feature#L216)
- [coreApiFavorites/favorites.feature:217](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favorites.feature#L217)
- [coreApiFavorites/favorites.feature:222](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favorites.feature#L222)
- [coreApiFavorites/favoritesSharingToShares.feature:67](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favoritesSharingToShares.feature#L67)
- [coreApiFavorites/favoritesSharingToShares.feature:68](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favoritesSharingToShares.feature#L68)


#### [WWW-Authenticate header for unauthenticated requests is not clear](https://github.com/owncloud/ocis/issues/2285)

- [coreApiWebdavOperations/refuseAccess.feature:22](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/refuseAccess.feature#L22)
- [coreApiWebdavOperations/refuseAccess.feature:23](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/refuseAccess.feature#L23)

#### [App Passwords/Tokens for legacy WebDAV clients](https://github.com/owncloud/ocis/issues/197)

- [coreApiAuthWebDav/webDavDELETEAuth.feature:135](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavDELETEAuth.feature#L135)
- [coreApiAuthWebDav/webDavDELETEAuth.feature:149](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavDELETEAuth.feature#L149)
- [coreApiAuthWebDav/webDavDELETEAuth.feature:161](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavDELETEAuth.feature#L161)
- [coreApiAuthWebDav/webDavDELETEAuth.feature:175](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthWebDav/webDavDELETEAuth.feature#L175)

#### [Request to edit non-existing user by authorized admin gets unauthorized in http response](https://github.com/owncloud/core/issues/38423)

- [coreApiAuthOcs/ocsPUTAuth.feature:24](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiAuthOcs/ocsPUTAuth.feature#L24)

#### [Sharing a same file twice to the same group](https://github.com/owncloud/ocis/issues/1710)

- [coreApiShareManagementBasicToShares/createShareToSharesFolder.feature:730](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/createShareToSharesFolder.feature#L730)
- [coreApiShareManagementBasicToShares/createShareToSharesFolder.feature:731](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/createShareToSharesFolder.feature#L731)

#### [PATCH request for TUS upload with wrong checksum gives incorrect response](https://github.com/owncloud/ocis/issues/1755)

- [coreApiWebdavUploadTUS/checksums.feature:84](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L84)
- [coreApiWebdavUploadTUS/checksums.feature:85](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L85)
- [coreApiWebdavUploadTUS/checksums.feature:86](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L86)
- [coreApiWebdavUploadTUS/checksums.feature:87](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L87)
- [coreApiWebdavUploadTUS/checksums.feature:92](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L92)
- [coreApiWebdavUploadTUS/checksums.feature:93](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L93)
- [coreApiWebdavUploadTUS/checksums.feature:173](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L173)
- [coreApiWebdavUploadTUS/checksums.feature:174](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L174)
- [coreApiWebdavUploadTUS/checksums.feature:179](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L179)
- [coreApiWebdavUploadTUS/checksums.feature:226](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L226)
- [coreApiWebdavUploadTUS/checksums.feature:227](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L227)
- [coreApiWebdavUploadTUS/checksums.feature:228](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L228)
- [coreApiWebdavUploadTUS/checksums.feature:229](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L229)
- [coreApiWebdavUploadTUS/checksums.feature:234](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L234)
- [coreApiWebdavUploadTUS/checksums.feature:235](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L235)
- [coreApiWebdavUploadTUS/checksums.feature:282](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L282)
- [coreApiWebdavUploadTUS/checksums.feature:283](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L283)
- [coreApiWebdavUploadTUS/checksums.feature:284](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L284)
- [coreApiWebdavUploadTUS/checksums.feature:285](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L285)
- [coreApiWebdavUploadTUS/checksums.feature:290](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L290)
- [coreApiWebdavUploadTUS/checksums.feature:291](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/checksums.feature#L291)
- [coreApiWebdavUploadTUS/optionsRequest.feature:8](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/optionsRequest.feature#L8)
- [coreApiWebdavUploadTUS/optionsRequest.feature:22](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/optionsRequest.feature#L22)
- [coreApiWebdavUploadTUS/optionsRequest.feature:35](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/optionsRequest.feature#L35)
- [coreApiWebdavUploadTUS/optionsRequest.feature:49](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/optionsRequest.feature#L49)
- [coreApiWebdavUploadTUS/uploadToShare.feature:176](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/uploadToShare.feature#L176)
- [coreApiWebdavUploadTUS/uploadToShare.feature:177](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/uploadToShare.feature#L177)
- [coreApiWebdavUploadTUS/uploadToShare.feature:195](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/uploadToShare.feature#L195)
- [coreApiWebdavUploadTUS/uploadToShare.feature:196](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/uploadToShare.feature#L196)
- [coreApiWebdavUploadTUS/uploadToShare.feature:214](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/uploadToShare.feature#L214)
- [coreApiWebdavUploadTUS/uploadToShare.feature:215](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/uploadToShare.feature#L215)
- [coreApiWebdavUploadTUS/uploadToShare.feature:253](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/uploadToShare.feature#L253)
- [coreApiWebdavUploadTUS/uploadToShare.feature:254](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/uploadToShare.feature#L254)
- [coreApiWebdavUploadTUS/uploadToShare.feature:295](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/uploadToShare.feature#L295)
- [coreApiWebdavUploadTUS/uploadToShare.feature:296](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/uploadToShare.feature#L296)

#### [TUS OPTIONS requests do not reply with TUS headers when invalid password](https://github.com/owncloud/ocis/issues/1012)

- [coreApiWebdavUploadTUS/optionsRequest.feature:62](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/optionsRequest.feature#L62)
- [coreApiWebdavUploadTUS/optionsRequest.feature:76](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/optionsRequest.feature#L76)
- [coreApiWebdavUploadTUS/optionsRequest.feature:89](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/optionsRequest.feature#L89)
- [coreApiWebdavUploadTUS/optionsRequest.feature:104](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUploadTUS/optionsRequest.feature#L104)

#### [Trying to accept a share with invalid ID gives incorrect OCS and HTTP status](https://github.com/owncloud/ocis/issues/2111)

- [coreApiShareOperationsToShares2/shareAccessByID.feature:85](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L85)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:86](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L86)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:87](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L87)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:88](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L88)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:89](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L89)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:90](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L90)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:91](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L91)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:92](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L92)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:104](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L104)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:105](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L105)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:136](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L136)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:137](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L137)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:138](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L138)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:139](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L139)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:140](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L140)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:141](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L141)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:142](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L142)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:143](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L143)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:155](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L155)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:156](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L156)


#### [Shares to deleted group listed in the response](https://github.com/owncloud/ocis/issues/2441)

- [coreApiShareManagementBasicToShares/createShareToSharesFolder.feature:534](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/createShareToSharesFolder.feature#L534)
- [coreApiShareManagementBasicToShares/createShareToSharesFolder.feature:535](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/createShareToSharesFolder.feature#L535)

#### [copying the file inside Shares folder returns 404](https://github.com/owncloud/ocis/issues/3874)

- [coreApiWebdavProperties1/copyFile.feature:410](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/copyFile.feature#L410)
- [coreApiWebdavProperties1/copyFile.feature:411](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/copyFile.feature#L411)
- [coreApiWebdavProperties1/copyFile.feature:416](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/copyFile.feature#L416)
- [coreApiWebdavProperties1/copyFile.feature:436](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/copyFile.feature#L436)
- [coreApiWebdavProperties1/copyFile.feature:437](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/copyFile.feature#L437)
- [coreApiWebdavProperties1/copyFile.feature:442](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/copyFile.feature#L442)

### Won't fix

Not everything needs to be implemented for ocis. While the oc10 testsuite covers these things we are not looking at them right now.

- _The `OC-LazyOps` header is [no longer supported by the client](https://github.com/owncloud/client/pull/8398), implementing this is not necessary for a first production release. We plan to have an upload state machine to visualize the state of a file, see https://github.com/owncloud/ocis/issues/214_
- _Blacklisted ignored files are no longer required because ocis can handle `.htaccess` files without security implications introduced by serving user provided files with apache._

#### [uploading with old-chunking does not work](https://github.com/owncloud/ocis/issues/1343)

- [coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature:20](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature#L20)
- [coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature:21](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature#L21)
- [coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature:26](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature#L26)
- [coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature:39](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature#L39)
- [coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature:40](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature#L40)
- [coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature:45](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature#L45)
- [coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature:80](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature#L80)
- [coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature:81](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature#L81)
- [coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature:86](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFileToExcludedDirectory.feature#L86)

#### [Blacklist files extensions](https://github.com/owncloud/ocis/issues/2177)

- [coreApiWebdavProperties1/copyFile.feature:120](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/copyFile.feature#L120)
- [coreApiWebdavProperties1/copyFile.feature:121](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/copyFile.feature#L121)
- [coreApiWebdavProperties1/copyFile.feature:126](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/copyFile.feature#L126)
- [coreApiWebdavProperties1/createFileFolder.feature:98](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/createFileFolder.feature#L98)
- [coreApiWebdavProperties1/createFileFolder.feature:99](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/createFileFolder.feature#L99)
- [coreApiWebdavProperties1/createFileFolder.feature:104](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavProperties1/createFileFolder.feature#L104)
- [coreApiWebdavUpload1/uploadFile.feature:181](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFile.feature#L181)
- [coreApiWebdavUpload1/uploadFile.feature:182](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFile.feature#L182)
- [coreApiWebdavUpload1/uploadFile.feature:187](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFile.feature#L187)
- [coreApiWebdavUpload2/uploadFileToBlacklistedNameUsingOldChunking.feature:19](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload2/uploadFileToBlacklistedNameUsingOldChunking.feature#L19)
- [coreApiWebdavUpload2/uploadFileToBlacklistedNameUsingOldChunking.feature:35](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload2/uploadFileToBlacklistedNameUsingOldChunking.feature#L35)
- [coreApiWebdavUpload2/uploadFileToBlacklistedNameUsingOldChunking.feature:36](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload2/uploadFileToBlacklistedNameUsingOldChunking.feature#L36)
- [coreApiWebdavUpload2/uploadFileToBlacklistedNameUsingOldChunking.feature:37](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload2/uploadFileToBlacklistedNameUsingOldChunking.feature#L37)
- [coreApiWebdavUpload2/uploadFileToExcludedDirectoryUsingOldChunking.feature:13](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload2/uploadFileToExcludedDirectoryUsingOldChunking.feature#L13)
- [coreApiWebdavUpload2/uploadFileToExcludedDirectoryUsingOldChunking.feature:20](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload2/uploadFileToExcludedDirectoryUsingOldChunking.feature#L20)
- [coreApiWebdavUpload2/uploadFileToExcludedDirectoryUsingOldChunking.feature:38](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload2/uploadFileToExcludedDirectoryUsingOldChunking.feature#L38)
- [coreApiWebdavUpload2/uploadFileToExcludedDirectoryUsingOldChunking.feature:39](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload2/uploadFileToExcludedDirectoryUsingOldChunking.feature#L39)
- [coreApiWebdavUpload2/uploadFileToExcludedDirectoryUsingOldChunking.feature:40](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload2/uploadFileToExcludedDirectoryUsingOldChunking.feature#L40)
- [coreApiWebdavMove2/moveFile.feature:287](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFile.feature#L287)
- [coreApiWebdavMove2/moveFile.feature:288](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFile.feature#L288)
- [coreApiWebdavMove2/moveFile.feature:293](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFile.feature#L293)

#### [cannot set blacklisted file names](https://github.com/owncloud/product/issues/260)

- [coreApiWebdavMove1/moveFolderToBlacklistedName.feature:21](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToBlacklistedName.feature#L21)
- [coreApiWebdavMove1/moveFolderToBlacklistedName.feature:22](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToBlacklistedName.feature#L22)
- [coreApiWebdavMove1/moveFolderToBlacklistedName.feature:27](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToBlacklistedName.feature#L27)
- [coreApiWebdavMove1/moveFolderToBlacklistedName.feature:40](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToBlacklistedName.feature#L40)
- [coreApiWebdavMove1/moveFolderToBlacklistedName.feature:41](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToBlacklistedName.feature#L41)
- [coreApiWebdavMove1/moveFolderToBlacklistedName.feature:46](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToBlacklistedName.feature#L46)
- [coreApiWebdavMove1/moveFolderToBlacklistedName.feature:81](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToBlacklistedName.feature#L81)
- [coreApiWebdavMove1/moveFolderToBlacklistedName.feature:82](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToBlacklistedName.feature#L82)
- [coreApiWebdavMove1/moveFolderToBlacklistedName.feature:87](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToBlacklistedName.feature#L87)
- [coreApiWebdavMove2/moveFileToBlacklistedName.feature:19](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToBlacklistedName.feature#L19)
- [coreApiWebdavMove2/moveFileToBlacklistedName.feature:20](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToBlacklistedName.feature#L20)
- [coreApiWebdavMove2/moveFileToBlacklistedName.feature:35](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToBlacklistedName.feature#L35)
- [coreApiWebdavMove2/moveFileToBlacklistedName.feature:36](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToBlacklistedName.feature#L36)
- [coreApiWebdavMove2/moveFileToBlacklistedName.feature:74](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToBlacklistedName.feature#L74)
- [coreApiWebdavMove2/moveFileToBlacklistedName.feature:75](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToBlacklistedName.feature#L75)

#### [cannot set excluded directories](https://github.com/owncloud/product/issues/261)

- [coreApiWebdavMove1/moveFolderToExcludedDirectory.feature:22](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToExcludedDirectory.feature#L22)
- [coreApiWebdavMove1/moveFolderToExcludedDirectory.feature:23](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToExcludedDirectory.feature#L23)
- [coreApiWebdavMove1/moveFolderToExcludedDirectory.feature:28](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToExcludedDirectory.feature#L28)
- [coreApiWebdavMove1/moveFolderToExcludedDirectory.feature:42](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToExcludedDirectory.feature#L42)
- [coreApiWebdavMove1/moveFolderToExcludedDirectory.feature:43](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToExcludedDirectory.feature#L43)
- [coreApiWebdavMove1/moveFolderToExcludedDirectory.feature:48](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToExcludedDirectory.feature#L48)
- [coreApiWebdavMove1/moveFolderToExcludedDirectory.feature:84](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToExcludedDirectory.feature#L84)
- [coreApiWebdavMove1/moveFolderToExcludedDirectory.feature:85](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToExcludedDirectory.feature#L85)
- [coreApiWebdavMove1/moveFolderToExcludedDirectory.feature:90](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolderToExcludedDirectory.feature#L90)
- [coreApiWebdavMove2/moveFileToExcludedDirectory.feature:20](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToExcludedDirectory.feature#L20)
- [coreApiWebdavMove2/moveFileToExcludedDirectory.feature:21](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToExcludedDirectory.feature#L21)
- [coreApiWebdavMove2/moveFileToExcludedDirectory.feature:37](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToExcludedDirectory.feature#L37)
- [coreApiWebdavMove2/moveFileToExcludedDirectory.feature:38](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToExcludedDirectory.feature#L38)
- [coreApiWebdavMove2/moveFileToExcludedDirectory.feature:78](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToExcludedDirectory.feature#L78)
- [coreApiWebdavMove2/moveFileToExcludedDirectory.feature:79](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToExcludedDirectory.feature#L79)

#### [system configuration options missing](https://github.com/owncloud/ocis/issues/1323)

- [coreApiWebdavUpload1/uploadFileToBlacklistedName.feature:31](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFileToBlacklistedName.feature#L31)
- [coreApiWebdavUpload1/uploadFileToBlacklistedName.feature:32](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFileToBlacklistedName.feature#L32)
- [coreApiWebdavUpload1/uploadFileToBlacklistedName.feature:37](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFileToBlacklistedName.feature#L37)
- [coreApiWebdavUpload1/uploadFileToBlacklistedName.feature:71](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFileToBlacklistedName.feature#L71)
- [coreApiWebdavUpload1/uploadFileToBlacklistedName.feature:70](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFileToBlacklistedName.feature#L70)
- [coreApiWebdavUpload1/uploadFileToBlacklistedName.feature:76](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavUpload1/uploadFileToBlacklistedName.feature#L76)

#### [Allow public link sharing only for certain groups feature not implemented](https://github.com/owncloud/ocis/issues/4623)

- [coreApiSharePublicLink3/allowGroupToCreatePublicLinks.feature:35](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink3/allowGroupToCreatePublicLinks.feature#L35)
- [coreApiSharePublicLink3/allowGroupToCreatePublicLinks.feature:92](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiSharePublicLink3/allowGroupToCreatePublicLinks.feature#L92)

#### [Preview of text file with UTF content does not render correctly](https://github.com/owncloud/ocis/issues/2570)

- [coreApiWebdavPreviews/previews.feature:211](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavPreviews/previews.feature#L211)

#### [Share path in the response is different between share states](https://github.com/owncloud/ocis/issues/2540)

- [coreApiShareManagementToShares/acceptShares.feature:65](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/acceptShares.feature#L65)
- [coreApiShareManagementToShares/acceptShares.feature:88](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/acceptShares.feature#L88)
- [coreApiShareManagementToShares/acceptShares.feature:214](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/acceptShares.feature#L214)
- [coreApiShareManagementToShares/acceptShares.feature:237](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/acceptShares.feature#L237)
- [coreApiShareManagementToShares/acceptShares.feature:275](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/acceptShares.feature#L275)
- [coreApiShareManagementToShares/acceptShares.feature:309](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/acceptShares.feature#L309)
- [coreApiShareManagementToShares/acceptShares.feature:529](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/acceptShares.feature#L529)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:123](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L123)
- [coreApiShareOperationsToShares2/shareAccessByID.feature:124](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareOperationsToShares2/shareAccessByID.feature#L124)
- [coreApiShareManagementBasicToShares/deleteShareFromShares.feature:177](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/deleteShareFromShares.feature#L177)
- [coreApiShareManagementBasicToShares/deleteShareFromShares.feature:178](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/deleteShareFromShares.feature#L178)
- [coreApiShareManagementBasicToShares/deleteShareFromShares.feature:179](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/deleteShareFromShares.feature#L179)
- [coreApiShareManagementBasicToShares/deleteShareFromShares.feature:180](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/deleteShareFromShares.feature#L180)
- [coreApiShareManagementBasicToShares/deleteShareFromShares.feature:196](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/deleteShareFromShares.feature#L196)
- [coreApiShareManagementBasicToShares/deleteShareFromShares.feature:197](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/deleteShareFromShares.feature#L197)
- [coreApiShareManagementBasicToShares/deleteShareFromShares.feature:198](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/deleteShareFromShares.feature#L198)
- [coreApiShareManagementBasicToShares/deleteShareFromShares.feature:199](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementBasicToShares/deleteShareFromShares.feature#L199)

#### [Content-type is not multipart/byteranges when downloading file with Range Header](https://github.com/owncloud/ocis/issues/2677)

- [coreApiWebdavOperations/downloadFile.feature:229](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/downloadFile.feature#L229)
- [coreApiWebdavOperations/downloadFile.feature:230](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/downloadFile.feature#L230)
- [coreApiWebdavOperations/downloadFile.feature:235](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/downloadFile.feature#L235)

#### [moveShareInsideAnotherShare behaves differently on oCIS than oC10](https://github.com/owncloud/ocis/issues/3047)

- [coreApiShareManagementToShares/moveShareInsideAnotherShare.feature:25](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/moveShareInsideAnotherShare.feature#L25)
- [coreApiShareManagementToShares/moveShareInsideAnotherShare.feature:86](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/moveShareInsideAnotherShare.feature#L86)
- [coreApiShareManagementToShares/moveShareInsideAnotherShare.feature:100](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/moveShareInsideAnotherShare.feature#L100)

#### [Renaming resource to banned name is allowed in spaces webdav](https://github.com/owncloud/ocis/issues/3099)

- [coreApiWebdavMove1/moveFolder.feature:27](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolder.feature#L27)
- [coreApiWebdavMove1/moveFolder.feature:45](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolder.feature#L45)
- [coreApiWebdavMove1/moveFolder.feature:63](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove1/moveFolder.feature#L63)
- [coreApiWebdavMove2/moveFile.feature:224](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFile.feature#L224)
- [coreApiWebdavMove2/moveFileToBlacklistedName.feature:25](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToBlacklistedName.feature#L25)
- [coreApiWebdavMove2/moveFileToBlacklistedName.feature:41](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToBlacklistedName.feature#L41)
- [coreApiWebdavMove2/moveFileToBlacklistedName.feature:80](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToBlacklistedName.feature#L80)

#### [REPORT method on spaces returns an incorrect d:href response](https://github.com/owncloud/ocis/issues/3111)

- [coreApiFavorites/favorites.feature:121](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favorites.feature#L121)
- [coreApiFavorites/favorites.feature:148](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favorites.feature#L148)
- [coreApiFavorites/favorites.feature:270](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiFavorites/favorites.feature#L270)

#### [could not create system tag](https://github.com/owncloud/ocis/issues/3092)

- [coreApiWebdavOperations/search.feature:274](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/search.feature#L274)
- [coreApiWebdavOperations/search.feature:291](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/search.feature#L291)
- [coreApiWebdavOperations/search.feature:317](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/search.feature#L317)

#### [Cannot disable the dav propfind depth infinity for resources](https://github.com/owncloud/ocis/issues/3720)

- [coreApiWebdavOperations/listFiles.feature:383](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/listFiles.feature#L383)
- [coreApiWebdavOperations/listFiles.feature:384](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/listFiles.feature#L384)
- [coreApiWebdavOperations/listFiles.feature:389](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/listFiles.feature#L389)
- [coreApiWebdavOperations/listFiles.feature:408](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/listFiles.feature#L408)
- [coreApiWebdavOperations/listFiles.feature:413](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/listFiles.feature#L413)
- [coreApiWebdavOperations/listFiles.feature:427](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/listFiles.feature#L427)
- [coreApiWebdavOperations/listFiles.feature:428](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/listFiles.feature#L428)
- [coreApiWebdavOperations/listFiles.feature:433](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/listFiles.feature#L433)

#### [Renaming resource to excluded directory name is allowed in spaces webdav](https://github.com/owncloud/ocis/issues/3102)

- [coreApiWebdavMove2/moveFileToExcludedDirectory.feature:26](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToExcludedDirectory.feature#L26)
- [coreApiWebdavMove2/moveFileToExcludedDirectory.feature:43](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToExcludedDirectory.feature#L43)
- [coreApiWebdavMove2/moveFileToExcludedDirectory.feature:84](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavMove2/moveFileToExcludedDirectory.feature#L84)

### [graph/users: enable/disable users](https://github.com/owncloud/ocis/issues/3064)

- [coreApiWebdavOperations/refuseAccess.feature:35](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/refuseAccess.feature#L35)
- [coreApiWebdavOperations/refuseAccess.feature:36](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/refuseAccess.feature#L36)
- [coreApiWebdavOperations/refuseAccess.feature:41](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiWebdavOperations/refuseAccess.feature#L41)

#### [OCS status code zero](https://github.com/owncloud/ocis/issues/3621)

- [coreApiShareManagementToShares/moveReceivedShare.feature:32](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiShareManagementToShares/moveReceivedShare.feature#L32)

#### [HTTP status code differ while deleting file of another user's trash bin](https://github.com/owncloud/ocis/issues/3544)

- [coreApiTrashbin/trashbinDelete.feature:109](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinDelete.feature#L109)

#### [Problem accessing trashbin with personal space id](https://github.com/owncloud/ocis/issues/3639)

- [coreApiTrashbin/trashbinDelete.feature:35](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinDelete.feature#L35)
- [coreApiTrashbin/trashbinDelete.feature:36](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinDelete.feature#L36)
- [coreApiTrashbin/trashbinDelete.feature:58](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinDelete.feature#L58)
- [coreApiTrashbin/trashbinDelete.feature:85](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinDelete.feature#L85)
- [coreApiTrashbin/trashbinDelete.feature:132](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinDelete.feature#L132)
- [coreApiTrashbin/trashbinDelete.feature:155](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinDelete.feature#L155)
- [coreApiTrashbin/trashbinDelete.feature:180](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinDelete.feature#L180)
- [coreApiTrashbin/trashbinDelete.feature:205](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinDelete.feature#L205)
- [coreApiTrashbin/trashbinDelete.feature:242](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinDelete.feature#L242)
- [coreApiTrashbin/trashbinDelete.feature:279](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinDelete.feature#L279)
- [coreApiTrashbin/trashbinDelete.feature:328](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinDelete.feature#L328)
- [coreApiTrashbin/trashbinFilesFolders.feature:25](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L25)
- [coreApiTrashbin/trashbinFilesFolders.feature:41](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L41)
- [coreApiTrashbin/trashbinFilesFolders.feature:60](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L60)
- [coreApiTrashbin/trashbinFilesFolders.feature:81](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L81)
- [coreApiTrashbin/trashbinFilesFolders.feature:100](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L100)
- [coreApiTrashbin/trashbinFilesFolders.feature:136](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L136)
- [coreApiTrashbin/trashbinFilesFolders.feature:159](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L159)
- [coreApiTrashbin/trashbinFilesFolders.feature:339](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L339)
- [coreApiTrashbin/trashbinFilesFolders.feature:340](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L340)
- [coreApiTrashbin/trashbinFilesFolders.feature:341](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L341)
- [coreApiTrashbin/trashbinFilesFolders.feature:346](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L346)
- [coreApiTrashbin/trashbinFilesFolders.feature:347](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L347)
- [coreApiTrashbin/trashbinFilesFolders.feature:348](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L348)
- [coreApiTrashbin/trashbinFilesFolders.feature:349](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L349)
- [coreApiTrashbin/trashbinFilesFolders.feature:350](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L350)
- [coreApiTrashbin/trashbinFilesFolders.feature:351](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L351)
- [coreApiTrashbin/trashbinFilesFolders.feature:368](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L368)
- [coreApiTrashbin/trashbinFilesFolders.feature:388](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L388)
- [coreApiTrashbin/trashbinFilesFolders.feature:442](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L442)
- [coreApiTrashbin/trashbinFilesFolders.feature:479](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/coreApiTrashbin/trashbinFilesFolders.feature#L479)

Note: always have an empty line at the end of this file.
The bash script that processes this file requires that the last line has a newline on the end.