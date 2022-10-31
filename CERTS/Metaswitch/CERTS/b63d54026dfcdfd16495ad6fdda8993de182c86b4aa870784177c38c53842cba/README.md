# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Avid Communication SHAKEN Cert 742D

Tested At: 31 Oct 22 16:42 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 949 day(s)\
Subject: CN=Avid Communication SHAKEN Cert 742D, O=Avid Communication, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/1b0b321bad320b960a4f0fc8f2408d0daa110730

View: [Click to view](https://understandingwebpki.com/?cert=MIICXzCCAgWgAwIBAgIQSuZWNeUVoKuDY3EplT75YDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMDYwNzEyMjQ1MFoXDTI1MDYwNjEyMjQ1MFowWDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMEkF2aWQgQ29tbXVuaWNhdGlvbjEsMCoGA1UEAwwjQXZpZCBDb21tdW5pY2F0aW9uIFNIQUtFTiBDZXJ0IDc0MkQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARUk%2FXJ%2BBKKYIipM3npJm%2BUSkTDZEJ9EUi1EZf0a9DDFCoc7aAgtJq%2B7LNxbHxaw5WzJ24ExRs5v5xstnlRWg15o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ3NDJEMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFIIZivC8J5nnImQzeq%2FoPlYrCnB%2FMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCIQCVIQ0iOeqh6nxX9y4H4jBshFP7d7tGg233RNF54gU3LgIgGvGIqiFHBVh9BRfiU7SzZTCgx7895dxuokFmIKBoLCs%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_issuer_dn](../../ISSUES/e_sti_issuer_dn/README.md) | error | ATIS&#x2011;1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 742D' |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_sti_key_usage](../../ISSUES/e_sti_key_usage/README.md) | error | ATIS&#x2011;1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |


Generated: 31/10/2022 at 16:43:22