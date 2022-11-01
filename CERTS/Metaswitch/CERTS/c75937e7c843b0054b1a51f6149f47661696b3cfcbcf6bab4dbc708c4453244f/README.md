# STIR/SHAKEN CA Ecosystem Compliance

## Certificate USCellular SHAKEN Cert 6349

Tested At: 01 Nov 22 09:58 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 517 day(s)\
Subject: CN=USCellular SHAKEN Cert 6349, L=Ch, ST=IL, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://crs.sti.uscellular.com/certs/uscc_shaken_CA.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIICVDCCAfqgAwIBAgIQZxVLbPqJw5UINoYY2jqr1jAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDQwMTE1MjgwOFoXDTI0MDMzMTE1MjgwOFowTTELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAklMMQswCQYDVQQHDAJDaDEkMCIGA1UEAwwbVVNDZWxsdWxhciBTSEFLRU4gQ2VydCA2MzQ5MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEiEoISdkvzCjey078Gz%2FzNVnHeJvthSEgrK95WIXUG4yVRm%2FubkEiInh53VZlW27VHM02mWLmf5eIn8QOxmDtM6OB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYENjM0OTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBT1FRvpTRjqLKOVkjVOuFFf1kU3ojAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNIADBFAiEA%2BT8JXAe03JuPfBApV5psimLfX6WhkB3U2FRr0ubt7soCICZ8VWUSCjILY5ZS5sa3JZoIyh%2FiUPC2H0SO1UQPTC8p)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_subject](../../ISSUES/e_sti_subject/README.md) | error | ATIS&#x2011;1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_sti_key_usage](../../ISSUES/e_sti_key_usage/README.md) | error | ATIS&#x2011;1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_sti_issuer_dn](../../ISSUES/e_sti_issuer_dn/README.md) | error | ATIS&#x2011;1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject_cn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 01/11/2022 at 10:05:32