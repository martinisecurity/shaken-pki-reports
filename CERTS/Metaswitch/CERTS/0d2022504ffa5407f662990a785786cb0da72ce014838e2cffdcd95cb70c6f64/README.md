# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Hunter Communications Shaken Cert 660C

Tested At: 30 Nov 22 17:24 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 516 day(s)\
Subject: CN=Hunter Communications Shaken Cert 660C, O=Hunter Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/7fba9a7fc8b3131f6fcea50e668939fd26bbd4a3

[View certificate details](https://understandingwebpki.com/?cert=MIICZTCCAgugAwIBAgIQfVPt3WU6F26qGtt5qesiYDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDQzMDE3MDUxMFoXDTI0MDQyOTE3MDUxMFowXjELMAkGA1UEBhMCVVMxHjAcBgNVBAoMFUh1bnRlciBDb21tdW5pY2F0aW9uczEvMC0GA1UEAwwmSHVudGVyIENvbW11bmljYXRpb25zIFNoYWtlbiBDZXJ0IDY2MEMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARM45tp3De0tSfhLpV5qWcJAr2ht1WfGxCYtpmz%2FBTHJQghRu%2BNCePgGnfJlVbwtLIbEjOzqg%2BUgjbiCTW4j8xko4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ2NjBDMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFEhp9NlJqOUGdNOISNUL%2FpuLqDKkMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIC395ONlYIT71KYVos5kRbg6H%2BdFI13gFiZRdqg8Rl1%2FAiEAiPSGuwiajPR9udIJYSP%2FlN3CFxjpf4i04Yq9GOPgIto%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 30 Nov 22 17:24 UTC