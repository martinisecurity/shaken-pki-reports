# STIR/SHAKEN CA Ecosystem Compliance

## Certificate CBTS Technology Solutions SHAKEN Cert 600F

Tested At: 17 Nov 22 19:20 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 515 day(s)\
Subject: CN=CBTS Technology Solutions SHAKEN Cert 600F, O=CBTS Technology Solutions LLC, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/de39b976c7549358da256530d81cfb1f446a2583

[View certificate details](https://understandingwebpki.com/?cert=MIICcDCCAhegAwIBAgIQEyjDL9%2BGClPgy36jYuaJEjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDQxNjE1MjcyNloXDTI0MDQxNTE1MjcyNlowajELMAkGA1UEBhMCVVMxJjAkBgNVBAoMHUNCVFMgVGVjaG5vbG9neSBTb2x1dGlvbnMgTExDMTMwMQYDVQQDDCpDQlRTIFRlY2hub2xvZ3kgU29sdXRpb25zIFNIQUtFTiBDZXJ0IDYwMEYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASZ48WyhCaIcoS7QbohiDbB2KxeWNW27aVarue%2FEDCYssE%2B8KzK3Px%2Bj3JnrHFgaNJNpP79mu0u1rCQaGCWb%2FbFo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ2MDBGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFGSxKaE7ej4Age7tofagLOz1Q4trMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIDNKt%2F54Bdz9o1RrrGEFGy%2FtZxr0SgG6tMy7OFc7%2FftBAiAfpnKqE%2FucbaQ5w9mtOMATmovpdHZgPNpQIDUv6XBtZA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 17 Nov 22 19:21 UTC