# STIR/SHAKEN CA Ecosystem Compliance
## Metaswitch

### Certificate 2299a3e38166a08cba68879b5717bec1c6434547
Tested At: 2022-10-28 16:27:03 +0000 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 415 day(s)\
Subject: CN=TDS Telecom SHAKEN Cert 7804, O=TDS Telecom, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1

Link: https://sti-cr.cgah.tnsi.com/certs/9b2b8bbbd5a5a28082425a21229434097aafd9b2

View: [Click to view](https://understandingwebpki.com/?cert=MIICUDCCAfegAwIBAgIQAjWjtyylnMU8unkpbpEwzzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIwMTIxNzE1MTkzN1oXDTIzMTIxNzE1MTkzN1owSjELMAkGA1UEBhMCVVMxFDASBgNVBAoMC1REUyBUZWxlY29tMSUwIwYDVQQDDBxURFMgVGVsZWNvbSBTSEFLRU4gQ2VydCA3ODA0MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEEzEaJxV7d0lzjaYSm21rM8ANhfLgaXBHx3jrSFzQdB%2BlKw4jpas6MZs3rur3jl65g%2Bz%2Fy1WKXRgzVZ3Ue3m2mKOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYENzgwNDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBSqrTbAPlsHb0WlVYwqpG90QLPgujAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNHADBEAiBE6jsXL3xjQuSogiDC%2BP2CN28wCpK%2F8FLchPBEuuCTDQIgPa4EfVwh9tifoJi1TQs%2BkpM%2FxSOmhk5PIJCqiOezYzU%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_key_usage | error | ATIS-1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| e_sti_issuer_dn | error | ATIS-1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |

### Not Effective

- e_sti_subject_cn
- e_cp1_3_ambiguous_identifier
- e_sti_serial_number
- w_cp_1_3_subject_email
- e_cp1_3_subject_sn
- e_sti_signature_algorithm
- e_sti_extension_unknown

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 28/10/2022 at 16:28:22