# STIR/SHAKEN CA Ecosystem Compliance

## Certificate CBTS Technology Solutions SHAKEN Cert 600F

Tested At: 24 Nov 23 11:00 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 144 day(s)\
Subject: CN=CBTS Technology Solutions SHAKEN Cert 600F, O=CBTS Technology Solutions LLC, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/1fa17ac42885f55c6be50d2eb8b93b2a88085fa9

[View certificate details](https://understandingwebpki.com/?cert=MIICcDCCAhegAwIBAgIQEyjDL9%2BGClPgy36jYuaJEjAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDQxNjE1MjcyNloXDTI0MDQxNTE1MjcyNlowajELMAkGA1UEBhMCVVMxJjAkBgNVBAoMHUNCVFMgVGVjaG5vbG9neSBTb2x1dGlvbnMgTExDMTMwMQYDVQQDDCpDQlRTIFRlY2hub2xvZ3kgU29sdXRpb25zIFNIQUtFTiBDZXJ0IDYwMEYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASZ48WyhCaIcoS7QbohiDbB2KxeWNW27aVarue%2FEDCYssE%2B8KzK3Px%2Bj3JnrHFgaNJNpP79mu0u1rCQaGCWb%2FbFo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ2MDBGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFGSxKaE7ej4Age7tofagLOz1Q4trMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIDNKt%2F54Bdz9o1RrrGEFGy%2FtZxr0SgG6tMy7OFc7%2FftBAiAfpnKqE%2FucbaQ5w9mtOMATmovpdHZgPNpQIDUv6XBtZA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_crl_distribution_struct
- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 24 Nov 23 11:17 UTC