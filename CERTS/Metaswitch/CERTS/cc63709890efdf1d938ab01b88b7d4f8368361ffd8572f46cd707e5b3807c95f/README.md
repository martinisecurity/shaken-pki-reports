# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Encore Communications SHAKEN Cert 956H

Tested At: 15 Nov 23 15:51 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 625 day(s)\
Subject: CN=Encore Communications SHAKEN Cert 956H, O=Encore Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/a1b25708d639f93a043a23170376e6bc24aea56c

[View certificate details](https://understandingwebpki.com/?cert=MIICZDCCAgugAwIBAgIQIK4LjinBRMxFLWYhmrJYljAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMDgwMjEwNDMwM1oXDTI1MDgwMTEwNDMwM1owXjELMAkGA1UEBhMCVVMxHjAcBgNVBAoMFUVuY29yZSBDb21tdW5pY2F0aW9uczEvMC0GA1UEAwwmRW5jb3JlIENvbW11bmljYXRpb25zIFNIQUtFTiBDZXJ0IDk1NkgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARI74cjIb%2BrPqf6GBpi8NPczFTMUT0hN9uOQjuWf85xGNoIdRyZYcoIOHqWFv1Io0N%2FeR2H%2BmlzX4WN4jQ57FzHo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ5NTZIMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFF7kzE3d%2Bb3f1pZmSf9uIGaklFHDMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIAkxulNHgPzT%2Bo1qVduGj69gerGZHLl7aH7L9fs7f533AiARZFnht0V0pR8k119FfEoA9p5pTkrbpBfLvGIVvypXaQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 956H' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 15 Nov 23 16:51 UTC