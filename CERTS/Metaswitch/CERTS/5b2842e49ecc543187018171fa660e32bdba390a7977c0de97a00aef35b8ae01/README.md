# STIR/SHAKEN CA Ecosystem Compliance

## Certificate New Horizon SHAKEN Cert 127E

Tested At: 06 Jul 23 13:53 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 380 day(s)\
Subject: CN=New Horizon SHAKEN Cert 127E, O=New Horizon Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/934dd1fd3cc80f18d4322f0336a0dab2e9d7998d

[View certificate details](https://understandingwebpki.com/?cert=MIICYTCCAgagAwIBAgIQGmUh5W%2FqbEKpk1xM4G1FFDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDcyMDE5NTQzM1oXDTI0MDcxOTE5NTQzM1owWTELMAkGA1UEBhMCVVMxIzAhBgNVBAoMGk5ldyBIb3Jpem9uIENvbW11bmljYXRpb25zMSUwIwYDVQQDDBxOZXcgSG9yaXpvbiBTSEFLRU4gQ2VydCAxMjdFMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEEwm9so1mzTd%2BAwu1bitn%2FJ463Rji6OiOFfX7JysJEp0fFypZnN09rXCBetv%2BKI6DYPgvkdnnR0D4LpxdpwWf7aOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMTI3RTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBRgpFF6d%2FmHAOqEElOH1YRoc%2BgYbjAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNJADBGAiEA92qHOLXPikuRVLQfsavW293v4V1wkkO14rf5S8RtXyUCIQDrK9qD34mofgxBk6cp5SpPt0%2FJQRoQrv%2FGBqNUK%2FDeeA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 06 Jul 23 14:08 UTC