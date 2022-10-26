# STIR/SHAKEN CA Ecosystem Compliance
## Neustar

### Certificate d0f263cc754caa6e444b9f3c350b78fb24f64250
Tested At: 2022-10-26 21:14:21 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 235 day(s)\
Subject: CN=SHAKEN 573J, O=Voyant Communications LLC, L=Chicago\\,ST=IL, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US

Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/49.162

View: [Click to view](https://understandingwebpki.com/?cert=MIIDITCCAsigAwIBAgIUdza%2FDUPf2ApqooLGiNthR31JBm4wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYxNzIxMjUwMloXDTIzMDYxNzIxMjUwMlowXzELMAkGA1UEBhMCVVMxFjAUBgNVBAcMDUNoaWNhZ28sU1Q9SUwxIjAgBgNVBAoMGVZveWFudCBDb21tdW5pY2F0aW9ucyBMTEMxFDASBgNVBAMMC1NIQUtFTiA1NzNKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE%2FURq0eJGdOQnZ0UXooXbXQZrDFd5fyooH84SvEOg%2FeWrYXnbSAPZzAoLmghkpMPvtJV4TV38k1glQJ22VVkPvqOCATkwggE1MBYGCCsGAQUFBwEaBAowCKAGFgQ1NzNKMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUr9HIwu5yTIP8P%2B0Zp20dkLIH8DowWwYIKwYBBQUHAQEETzBNMEsGCCsGAQUFBzAChj9odHRwOi8vY2FjZXJ0cy11cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydCAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAdBgNVHQ4EFgQU66222dMCEHzAuVp55rdOktMcUREwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCICkYJHTiqtj7eBlY0zHjCcIHjh35sNCm8bVO%2FG%2F4yZ6yAiByuXzCeyyF%2FNzmUgHmajahVSkWZkx9cZmTxgW%2FaCPaEw%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_extension_unknown | error | ATIS-1000080v4 | STI certificate shall not include extensions that are not specified |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |

Generated: 26/10/2022 at 21:14:23