# STIR/SHAKEN CA Ecosystem Compliance

## Certificate TDS Telecom SHAKEN Cert 7804

Tested At: 04 Oct 24 15:31 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: -292 day(s)\
Subject: CN=TDS Telecom SHAKEN Cert 7804, O=TDS Telecom, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/9ce9430dbe64eaf1eded88e8494d43dd11ee6ae0

[View certificate details](https://x509.io/?cert=MIICUDCCAfegAwIBAgIQAjWjtyylnMU8unkpbpEwzzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIwMTIxNzE1MTkzN1oXDTIzMTIxNzE1MTkzN1owSjELMAkGA1UEBhMCVVMxFDASBgNVBAoMC1REUyBUZWxlY29tMSUwIwYDVQQDDBxURFMgVGVsZWNvbSBTSEFLRU4gQ2VydCA3ODA0MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEEzEaJxV7d0lzjaYSm21rM8ANhfLgaXBHx3jrSFzQdB%2BlKw4jpas6MZs3rur3jl65g%2Bz%2Fy1WKXRgzVZ3Ue3m2mKOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYENzgwNDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBSqrTbAPlsHb0WlVYwqpG90QLPgujAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNHADBEAiBE6jsXL3xjQuSogiDC%2BP2CN28wCpK%2F8FLchPBEuuCTDQIgPa4EfVwh9tifoJi1TQs%2BkpM%2FxSOmhk5PIJCqiOezYzU%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
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


Generated: 04 Oct 24 16:29 UTC