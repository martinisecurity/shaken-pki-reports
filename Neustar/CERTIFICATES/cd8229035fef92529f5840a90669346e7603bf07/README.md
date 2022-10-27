# STIR/SHAKEN CA Ecosystem Compliance
## Neustar

### Certificate cd8229035fef92529f5840a90669346e7603bf07
Tested At: 2022-10-27 18:57:00 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 195 day(s)\
Subject: CN=SHAKEN 402E, O=IP Networked Services Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US

Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11304.10152

View: [Click to view](https://understandingwebpki.com/?cert=MIIDCTCCArCgAwIBAgIUWz%2BK%2BE%2FIQOvmgMQ5OKo9WDZoI8MwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDUxMDE4MTczMFoXDTIzMDUxMDE4MTczMFowRzELMAkGA1UEBhMCVVMxIjAgBgNVBAoMGUlQIE5ldHdvcmtlZCBTZXJ2aWNlcyBJbmMxFDASBgNVBAMMC1NIQUtFTiA0MDJFMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEy3%2BgUZVLgi7lrnMbPxH7I3Fh7L30DeZ2pZ4BHLaM4SZoldag85Ee1vSGp5IXVx1nyBj0BjOItXDDMlU82yLKvKOCATkwggE1MBYGCCsGAQUFBwEaBAowCKAGFgQ0MDJFMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUr9HIwu5yTIP8P%2B0Zp20dkLIH8DowWwYIKwYBBQUHAQEETzBNMEsGCCsGAQUFBzAChj9odHRwOi8vY2FjZXJ0cy11cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydCAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAdBgNVHQ4EFgQUFO76%2FHW5WSLfz3X%2FLWYxd0ZcDpowDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIEH%2BWUFTEtHnqX7Ebq%2BRvhANwWZBL7AGDd2igxfyFQEnAiAqe04CtMyk00nSOenf3x1GwDC%2BySz%2FeM%2BgiRwBZeEcgw%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_extension_unknown | error | ATIS-1000080v4 | STI certificate shall not include extensions that are not specified |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |

Generated: 27/10/2022 at 18:57:26