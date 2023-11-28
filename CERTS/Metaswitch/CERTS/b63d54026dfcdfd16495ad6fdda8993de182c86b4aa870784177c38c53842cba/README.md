# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Avid Communication SHAKEN Cert 742D

Tested At: 28 Nov 23 10:18 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 557 day(s)\
Subject: CN=Avid Communication SHAKEN Cert 742D, O=Avid Communication, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/1b0b321bad320b960a4f0fc8f2408d0daa110730

[View certificate details](https://understandingwebpki.com/?cert=MIICXzCCAgWgAwIBAgIQSuZWNeUVoKuDY3EplT75YDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMDYwNzEyMjQ1MFoXDTI1MDYwNjEyMjQ1MFowWDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMEkF2aWQgQ29tbXVuaWNhdGlvbjEsMCoGA1UEAwwjQXZpZCBDb21tdW5pY2F0aW9uIFNIQUtFTiBDZXJ0IDc0MkQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARUk%2FXJ%2BBKKYIipM3npJm%2BUSkTDZEJ9EUi1EZf0a9DDFCoc7aAgtJq%2B7LNxbHxaw5WzJ24ExRs5v5xstnlRWg15o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ3NDJEMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFIIZivC8J5nnImQzeq%2FoPlYrCnB%2FMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQCVIQ0iOeqh6nxX9y4H4jBshFP7d7tGg233RNF54gU3LgIgGvGIqiFHBVh9BRfiU7SzZTCgx7895dxuokFmIKBoLCs%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 742D', but common name is 'Avid Communication SHAKEN Cert 742D' |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 10:53 UTC