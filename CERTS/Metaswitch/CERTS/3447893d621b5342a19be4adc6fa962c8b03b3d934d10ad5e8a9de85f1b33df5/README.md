# STIR/SHAKEN CA Ecosystem Compliance

## Certificate South Central Rural Telecommunications Coop SHAKEN Cert 0418

Tested At: 18 Aug 25 20:05 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 546 day(s)\
Subject: CN=South Central Rural Telecommunications Coop SHAKEN Cert 0418, O=South Central Rural Telephone, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/448db51c0154c1b4aa55612c8779cede93963e92

[View certificate details](https://x509.io/?cert=MIIChDCCAimgAwIBAgIQJRtDsyusexPDnWCX6LOmqTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTI0MDIxNjEwMjQ0NFoXDTI3MDIxNTEwMjQ0NFowfDELMAkGA1UEBhMCVVMxJjAkBgNVBAoMHVNvdXRoIENlbnRyYWwgUnVyYWwgVGVsZXBob25lMUUwQwYDVQQDDDxTb3V0aCBDZW50cmFsIFJ1cmFsIFRlbGVjb21tdW5pY2F0aW9ucyBDb29wIFNIQUtFTiBDZXJ0IDA0MTgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASZrp42qHXDRNLuOP%2Fb96pNd%2FM3mxwurg8uV8SIe%2BJwEG1D6xdPJpARqQwSH7ksf7twp9g8EsOZZ9eRmYYhIyozo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQwNDE4MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFIk8jqxyGNfVMHHXreiSfUUKqVONMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0kAMEYCIQCSL0iSs9kWK2h1YrH5zFDxaHNwx0beY0WcsObr7S5rhgIhANMtfY9td85RO12SjeqrRgk10D0CRCpOwWdL0%2BKzQXW6)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0418', but common name is 'South Central Rural Telecommunications Coop SHAKEN Cert 0418' |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 18 Aug 25 21:13 UTC