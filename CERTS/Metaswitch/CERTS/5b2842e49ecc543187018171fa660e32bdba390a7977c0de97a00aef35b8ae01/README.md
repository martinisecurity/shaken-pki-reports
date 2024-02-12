# STIR/SHAKEN CA Ecosystem Compliance

## Certificate New Horizon SHAKEN Cert 127E

Tested At: 12 Feb 24 16:26 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 159 day(s)\
Subject: CN=New Horizon SHAKEN Cert 127E, O=New Horizon Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/0ecd34983b628fdc7c5a4fc3ab800205cd77c78e

[View certificate details](https://understandingwebpki.com/?cert=MIICYTCCAgagAwIBAgIQGmUh5W%2FqbEKpk1xM4G1FFDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDcyMDE5NTQzM1oXDTI0MDcxOTE5NTQzM1owWTELMAkGA1UEBhMCVVMxIzAhBgNVBAoMGk5ldyBIb3Jpem9uIENvbW11bmljYXRpb25zMSUwIwYDVQQDDBxOZXcgSG9yaXpvbiBTSEFLRU4gQ2VydCAxMjdFMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEEwm9so1mzTd%2BAwu1bitn%2FJ463Rji6OiOFfX7JysJEp0fFypZnN09rXCBetv%2BKI6DYPgvkdnnR0D4LpxdpwWf7aOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMTI3RTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBRgpFF6d%2FmHAOqEElOH1YRoc%2BgYbjAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNJADBGAiEA92qHOLXPikuRVLQfsavW293v4V1wkkO14rf5S8RtXyUCIQDrK9qD34mofgxBk6cp5SpPt0%2FJQRoQrv%2FGBqNUK%2FDeeA%3D%3D)

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


Generated: 12 Feb 24 17:02 UTC