# STIR/SHAKEN CA Ecosystem Compliance

## Certificate prod SHAKEN 811J

Tested At: 01 Nov 22 09:56 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 362 day(s)\
Subject: CN=prod SHAKEN 811J, O=Alianza, L=Pleasant Grove, ST=Utah, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://api.alianza.com/v2/stir-shaken/certs/b45a4083-1554-4412-b5fc-bbd2c027091e/key.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDKDCCAs6gAwIBAgIUG8XRNt094oNulc2ltTa0lzH9aPgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTIyMTAyODE0NTkzNVoXDTIzMTAyODE0NTkzNVowYjELMAkGA1UEBhMCVVMxDTALBgNVBAgMBFV0YWgxFzAVBgNVBAcMDlBsZWFzYW50IEdyb3ZlMRAwDgYDVQQKDAdBbGlhbnphMRkwFwYDVQQDDBBwcm9kIFNIQUtFTiA4MTFKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAElnn93NPvWDVuPqwD8TmCh9ozF4j6OCbTcFrObVZ6DZ%2FS%2FWcodSyA%2BkugUN1PYWrP2M%2FwzgFajAgJYdNGMcrBW6OCATwwggE4MBYGCCsGAQUFBwEaBAowCKAGFgQ4MTFKMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUgk4V%2F%2F6famdR5MiXx210w%2FxlRXgwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAdBgNVHQ4EFgQU5zlnG7MzIXeihn%2FxrA8I8PhsiLYwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIFLhGa2FWXM2ouiKNu9rBX9HXpTBRZJTHhyw2pCc9VtpAiEA7oAGx8d7av2XexFwEoN3ixk8pxZ7XQfkf%2FcRTPgfFSU%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 01/11/2022 at 10:05:32