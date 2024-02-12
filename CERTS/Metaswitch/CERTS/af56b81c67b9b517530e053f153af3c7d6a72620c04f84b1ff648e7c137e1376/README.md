# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Fidelity Communications SHAKEN Cert 1882

Tested At: 12 Feb 24 16:26 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 965 day(s)\
Subject: CN=Fidelity Communications SHAKEN Cert 1882, O=Fidelity Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/514fb1f7a1891a792528e0cf023b424c55ddb2cd

[View certificate details](https://understandingwebpki.com/?cert=MIICajCCAg%2BgAwIBAgIQWFh3D8v7dLUHHF5RvwsAaDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMTAwNTEzNDQxNloXDTI2MTAwNDEzNDQxNlowYjELMAkGA1UEBhMCVVMxIDAeBgNVBAoMF0ZpZGVsaXR5IENvbW11bmljYXRpb25zMTEwLwYDVQQDDChGaWRlbGl0eSBDb21tdW5pY2F0aW9ucyBTSEFLRU4gQ2VydCAxODgyMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEpHlCsPJXu153scWGBxdHRxQjmGEju0k9yVxBQQxkNUCvh8DlQJf1XkCxkdaQqdw5eKG4xaOMXZMOJ4GeHa14zKOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMTg4MjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRqD6ejLc5WSYF9qWncIxOFDjya1jAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNJADBGAiEA5GqZDCDfT6vs4PDl9MggW1RmZ%2FjApqwrm0RiEjzpAYMCIQDgO1oE9zWegdS65lRUeLR0F%2FicsxoFNziAzz%2FhBqoweg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 1882', but common name is 'Fidelity Communications SHAKEN Cert 1882' |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 12 Feb 24 17:02 UTC