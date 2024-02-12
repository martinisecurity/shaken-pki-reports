# STIR/SHAKEN CA Ecosystem Compliance

## Certificate EPB Telecom SHAKEN Cert 4645

Tested At: 12 Feb 24 16:26 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 857 day(s)\
Subject: CN=EPB Telecom SHAKEN Cert 4645, O=Electric Power Board of Chattanooga dba EPB Telecom, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/245819f978821d3e211a6956b9e0d09fd3fed483

[View certificate details](https://understandingwebpki.com/?cert=MIICeTCCAh%2BgAwIBAgIQCEIWWvsZXDdiKjxAz3JWajAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDYxOTA5MjI1OVoXDTI2MDYxODA5MjI1OVowcjELMAkGA1UEBhMCVVMxPDA6BgNVBAoMM0VsZWN0cmljIFBvd2VyIEJvYXJkIG9mIENoYXR0YW5vb2dhIGRiYSBFUEIgVGVsZWNvbTElMCMGA1UEAwwcRVBCIFRlbGVjb20gU0hBS0VOIENlcnQgNDY0NTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPGR8wRm0GHjBNsULGMgt0REIfK64%2BCrjd4u%2BlY25p71Kvaa3JMhw%2BuMJcZ9lOWI4IxnYqxQ1MGJ4Blez0cRNUajgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDQ2NDUwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUWhoYrxsbk%2F2gF%2FCe3iyGjSGQZxgwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIhAKi7gj8XlyvgCXNjg4GSxHlxlKWTun1xAIXJf5d%2BTmEdAiBBu%2FBHw%2BGc2JKw5g6ooIMG5EYVeKG4Nla%2BI1%2FXLv0oVA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 4645', but common name is 'EPB Telecom SHAKEN Cert 4645' |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |


Generated: 12 Feb 24 17:02 UTC