# STIR/SHAKEN CA Ecosystem Compliance

## Certificate USCellular SHAKEN Cert 6349

Tested At: 21 Nov 23 17:08 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 131 day(s)\
Subject: CN=USCellular SHAKEN Cert 6349, L=Ch, ST=IL, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://robocall.sti.uscellular.com/certs/uscc_shaken_CA.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICVDCCAfqgAwIBAgIQZxVLbPqJw5UINoYY2jqr1jAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDQwMTE1MjgwOFoXDTI0MDMzMTE1MjgwOFowTTELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAklMMQswCQYDVQQHDAJDaDEkMCIGA1UEAwwbVVNDZWxsdWxhciBTSEFLRU4gQ2VydCA2MzQ5MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEiEoISdkvzCjey078Gz%2FzNVnHeJvthSEgrK95WIXUG4yVRm%2FubkEiInh53VZlW27VHM02mWLmf5eIn8QOxmDtM6OB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYENjM0OTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBT1FRvpTRjqLKOVkjVOuFFf1kU3ojAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNIADBFAiEA%2BT8JXAe03JuPfBApV5psimLfX6WhkB3U2FRr0ubt7soCICZ8VWUSCjILY5ZS5sa3JZoIyh%2FiUPC2H0SO1UQPTC8p)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_atis_subject](../../ISSUES/e_atis_subject/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:16 UTC