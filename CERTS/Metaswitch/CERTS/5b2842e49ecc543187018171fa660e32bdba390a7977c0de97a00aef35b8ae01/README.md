# STIR/SHAKEN CA Ecosystem Compliance

## Certificate New Horizon SHAKEN Cert 127E

Tested At: 31 Oct 22 16:42 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 628 day(s)\
Subject: CN=New Horizon SHAKEN Cert 127E, O=New Horizon Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/934dd1fd3cc80f18d4322f0336a0dab2e9d7998d

View: [Click to view](https://understandingwebpki.com/?cert=MIICYTCCAgagAwIBAgIQGmUh5W%2FqbEKpk1xM4G1FFDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDcyMDE5NTQzM1oXDTI0MDcxOTE5NTQzM1owWTELMAkGA1UEBhMCVVMxIzAhBgNVBAoMGk5ldyBIb3Jpem9uIENvbW11bmljYXRpb25zMSUwIwYDVQQDDBxOZXcgSG9yaXpvbiBTSEFLRU4gQ2VydCAxMjdFMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEEwm9so1mzTd%2BAwu1bitn%2FJ463Rji6OiOFfX7JysJEp0fFypZnN09rXCBetv%2BKI6DYPgvkdnnR0D4LpxdpwWf7aOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMTI3RTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBRgpFF6d%2FmHAOqEElOH1YRoc%2BgYbjAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNJADBGAiEA92qHOLXPikuRVLQfsavW293v4V1wkkO14rf5S8RtXyUCIQDrK9qD34mofgxBk6cp5SpPt0%2FJQRoQrv%2FGBqNUK%2FDeeA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_issuer_dn](../../ISSUES/e_sti_issuer_dn/README.md) | error | ATIS&#x2011;1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_sti_key_usage](../../ISSUES/e_sti_key_usage/README.md) | error | ATIS&#x2011;1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject_cn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 16:43:22