# STIR/SHAKEN CA Ecosystem Compliance

## Certificate USCellular SHAKEN Cert 6349

Tested At: 26 Aug 24 17:57 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 931 day(s)\
Subject: CN=USCellular SHAKEN Cert 6349, OU=QCall, O=USCC, L=Schaumburg, ST=Illinois, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://robocall.sti.uscellular.com/certs/uscellular2024certchain.crt

[View certificate details](https://x509.io/?cert=MIICgDCCAiegAwIBAgIQGCdyzrMiP16MGTaTa77HhTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTI0MDMxNTEwMjcwNloXDTI3MDMxNTEwMjcwNlowejELMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMRMwEQYDVQQHDApTY2hhdW1idXJnMQ0wCwYDVQQKDARVU0NDMQ4wDAYDVQQLDAVRQ2FsbDEkMCIGA1UEAwwbVVNDZWxsdWxhciBTSEFLRU4gQ2VydCA2MzQ5MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEPCz2GG2j0BaG8geGTt%2BDHEWDm7E%2Bi4oszTQ3j4OD3YtjBisr%2FoxkaKXZbPoJR%2BUtHCpQnl%2BMKYAtHjMmcNm6p6OB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYENjM0OTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBTQnyhUTci7UcvKVMZAigFXPTJE3DAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNHADBEAiBOPgLOuXz1H%2B6Avyefi4UCxgbV3s9y3c%2FCYowiVD%2FIHgIgYdY4gKYgxokD4zagqZT1Pj4EnZbrMyTfc5GBtAbVcSU%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 6349', but common name is 'USCellular SHAKEN Cert 6349' |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 26 Aug 24 18:03 UTC