# STIR/SHAKEN CA Ecosystem Compliance

## Certificate EssexTel INC SHAKEN 692J

Tested At: 28 Apr 23 02:11 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 13 day(s)\
Subject: O=EssexTel INC SHAKEN 692J, CN=EssexTel INC SHAKEN 692J\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://netnumber-sti-cr.s3.amazonaws.com/certs/f4fc911b-5676-4cdc-9722-cbcae0cbb55f

[View certificate details](https://understandingwebpki.com/?cert=MIICyzCCAlCgAwIBAgIILdnPpI0Y2JIwCgYIKoZIzj0EAwIwgY4xMDAuBgNVBAMMJ05ldE51bWJlciBTSEFLRU4gUm9vdCBJbnRlcm1lZGlhdGUgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMB4XDTIzMDQxMDEyNTk1NloXDTIzMDUxMDEyNTk1NlowRjEhMB8GA1UEAwwYRXNzZXhUZWwgSU5DIFNIQUtFTiA2OTJKMSEwHwYDVQQKDBhFc3NleFRlbCBJTkMgU0hBS0VOIDY5MkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ5IpHdHAXUnFrSqSla3wZzKKWhgTC1qtKD8nn4HiCTMOLYCZ9tvn%2Ff64BC33i%2FBvtWVPI99FnmzuijbuY%2Fszx2o4HeMIHbMBYGCCsGAQUFBwEaBAowCKAGFgQ2OTJKMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB8GA1UdIwQYMBaAFHEvyILcqOAi49%2BgaPn4XlyE3uW9MB0GA1UdDgQWBBQCBLAbXeef4FoyQ%2F%2BUD7Kz0MFdkzBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwGgYDVR0gAQH%2FBBAwDjAMBgpghkgBhv8JAQEBMAoGCCqGSM49BAMCA2kAMGYCMQCng%2B1e0Yi1PWd5z6QmYmg9w6Zb%2BvM1v0ukDJPDn0yBetTqZWbUJRGYMNJo7NaglSECMQCsYriOaZv%2BJAjaYI5hK%2BardLp3M0JS%2FsLfguO2IRI0xtbHefo8Pdp145LEvUSOSQU%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject](../../ISSUES/e_atis_subject/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [n_atis_certificate_policy_critical](../../ISSUES/n_atis_certificate_policy_critical/README.md) | notice | ATIS1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |


Generated: 28 Apr 23 02:17 UTC