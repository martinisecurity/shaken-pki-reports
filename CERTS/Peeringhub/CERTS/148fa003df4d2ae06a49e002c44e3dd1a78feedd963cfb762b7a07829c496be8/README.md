# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Definitely SHAKEN 709J but not SHAKEN 0007

Tested At: 01 Nov 22 22:12 UTC\
Initial Validity Period: 1 day(s)\
Remaining Validity Period: 1 day(s)\
Subject: CN=Definitely SHAKEN 709J but not SHAKEN 0007, OU=Trolling division, O=Low Latency Communications aka L.L.C. LLC, L=Birmingham Metro, ST=Alabama, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Link: https://sketchy.gay/shaken/llc-cert-2.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIDeDCCAx6gAwIBAgIQZlVVHs4j6EIEnkwYSfExuzAKBggqhkjOPQQDAjB8MQswCQYDVQQGEwJVUzEXMBUGA1UECgwOUGVlcmluZ2h1YiBJbmMxIjAgBgNVBAsMGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxMDAuBgNVBAMMJ1BlZXJpbmdodWIgSW5jIFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EgMjAeFw0yMjExMDExOTI2MzlaFw0yMjExMDIxODEzMzNaMIG%2FMQswCQYDVQQGEwJVUzEQMA4GA1UECAwHQWxhYmFtYTEZMBcGA1UEBwwQQmlybWluZ2hhbSBNZXRybzEyMDAGA1UECgwpTG93IExhdGVuY3kgQ29tbXVuaWNhdGlvbnMgYWthIEwuTC5DLiBMTEMxGjAYBgNVBAsMEVRyb2xsaW5nIGRpdmlzaW9uMTMwMQYDVQQDDCpEZWZpbml0ZWx5IFNIQUtFTiA3MDlKIGJ1dCBub3QgU0hBS0VOIDAwMDcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQUFr3dToowj31eDc%2FQ8BKysyAVPfiV6HuIFP57aofUf21dKDVMVjLwUAiHnqHdiZsJ8XeDFlM08R4lYkSMCzL9o4IBPDCCATgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFIavfwLKu%2BTJ2vKmMbOSV3KFLRGdMB8GA1UdIwQYMBaAFK6hc1GIKVcRygyp9LEKbk64S00HMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAWBggrBgEFBQcBGgQKMAigBhYENzA5SjCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwCgYIKoZIzj0EAwIDSAAwRQIhAPuvU26xzUw6V5kFSc1yeRh6tSEKY81J1X8hUuIzV%2FfSAiAaxd5qqz%2BCa9fjenve4JUgffdWCbanbniID3ghYf4eVw%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 01/11/2022 at 22:19:34