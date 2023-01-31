# STIR/SHAKEN CA Ecosystem Compliance

## Certificate USCellular SHAKEN Cert 6349

Tested At: 31 Jan 23 21:42 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 425 day(s)\
Subject: CN=USCellular SHAKEN Cert 6349, L=Ch, ST=IL, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://crs.sti.uscellular.com/certs/uscc_shaken_CA.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICVDCCAfqgAwIBAgIQZxVLbPqJw5UINoYY2jqr1jAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDQwMTE1MjgwOFoXDTI0MDMzMTE1MjgwOFowTTELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAklMMQswCQYDVQQHDAJDaDEkMCIGA1UEAwwbVVNDZWxsdWxhciBTSEFLRU4gQ2VydCA2MzQ5MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEiEoISdkvzCjey078Gz%2FzNVnHeJvthSEgrK95WIXUG4yVRm%2FubkEiInh53VZlW27VHM02mWLmf5eIn8QOxmDtM6OB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYENjM0OTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBT1FRvpTRjqLKOVkjVOuFFf1kU3ojAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNIADBFAiEA%2BT8JXAe03JuPfBApV5psimLfX6WhkB3U2FRr0ubt7soCICZ8VWUSCjILY5ZS5sa3JZoIyh%2FiUPC2H0SO1UQPTC8p)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject](../../ISSUES/e_atis_subject/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31 Jan 23 21:50 UTC