# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Crosstel Tandem Inc Holston Shaken Cert 308H 

Tested At: 18 Aug 25 20:05 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 263 day(s)\
Subject: CN=Crosstel Tandem Inc Holston Shaken Cert 308H\\ , O=Crosstel Tandem Inc Holston, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/2411390577fae283ff16682fc3c75cb0597b918e

[View certificate details](https://x509.io/?cert=MIICcTCCAhigAwIBAgIQX2Hejpz90T2%2FVt0qgFytXDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDUwOTEzNDIxN1oXDTI2MDUwODEzNDIxN1owazELMAkGA1UEBhMCVVMxJDAiBgNVBAoMG0Nyb3NzdGVsIFRhbmRlbSBJbmMgSG9sc3RvbjE2MDQGA1UEAwwtQ3Jvc3N0ZWwgVGFuZGVtIEluYyBIb2xzdG9uIFNoYWtlbiBDZXJ0IDMwOEggMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAENo71soA%2F1Xj98ROr4r0%2Bwn9k8bz5LyVLnYMxS7G3GcMN5YSODv3u%2BA%2FiVHtNEICxtlo%2FJfNx%2Bjcs%2BbavETZdkaOB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYEMzA4SDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBR5goKNnY8OWwOtHo2%2FBNzg0YtPaTAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNHADBEAiBXCMvTHMbYydAkzX%2BcfkoNeY%2F8SV40PN51WElBYaYM2gIgcM6L4wnRFdf7HOA5HPYXTYjLy3AXweK%2FVGomOMN4uHk%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 308H', but common name is 'Crosstel Tandem Inc Holston Shaken Cert 308H ' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'Crosstel Tandem Inc Holston Shaken Cert 308H ' does not contain 'SHAKEN' |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |


Generated: 18 Aug 25 21:13 UTC