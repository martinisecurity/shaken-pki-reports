# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Hunter Communications Shaken Cert 660C

Tested At: 12 Feb 24 16:26 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 78 day(s)\
Subject: CN=Hunter Communications Shaken Cert 660C, O=Hunter Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/7fba9a7fc8b3131f6fcea50e668939fd26bbd4a3

[View certificate details](https://understandingwebpki.com/?cert=MIICZTCCAgugAwIBAgIQfVPt3WU6F26qGtt5qesiYDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDQzMDE3MDUxMFoXDTI0MDQyOTE3MDUxMFowXjELMAkGA1UEBhMCVVMxHjAcBgNVBAoMFUh1bnRlciBDb21tdW5pY2F0aW9uczEvMC0GA1UEAwwmSHVudGVyIENvbW11bmljYXRpb25zIFNoYWtlbiBDZXJ0IDY2MEMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARM45tp3De0tSfhLpV5qWcJAr2ht1WfGxCYtpmz%2FBTHJQghRu%2BNCePgGnfJlVbwtLIbEjOzqg%2BUgjbiCTW4j8xko4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ2NjBDMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFEhp9NlJqOUGdNOISNUL%2FpuLqDKkMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIC395ONlYIT71KYVos5kRbg6H%2BdFI13gFiZRdqg8Rl1%2FAiEAiPSGuwiajPR9udIJYSP%2FlN3CFxjpf4i04Yq9GOPgIto%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'Hunter Communications Shaken Cert 660C' does not contain 'SHAKEN' |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Feb 24 17:02 UTC