# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Cspire SHAKEN Cert 6581

Tested At: 28 Nov 23 15:59 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 117 day(s)\
Subject: CN=Cspire SHAKEN Cert 6581, O=Cspire, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/2487f31aa2e09ad03fa91fa776793115d03392e9

[View certificate details](https://understandingwebpki.com/?cert=MIICRjCCAe2gAwIBAgIQE3lY2DnuwU0MPjNdy%2BIj3TAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDMyNDIzNTI0NFoXDTI0MDMyMzIzNTI0NFowQDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBkNzcGlyZTEgMB4GA1UEAwwXQ3NwaXJlIFNIQUtFTiBDZXJ0IDY1ODEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS1it%2Ffie%2BxEYwKEiByYLAolMSwz53x2oBUB9DOdK%2FIZqPk82h%2B8Ou1nvO%2FZlSmnsG5C81gqI4CHdX2w7nVlTtqo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ2NTgxMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFGjSBWqoiWThMqlMxnbXceiQcx2XMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIGuhzlPpon4tSFuUYrsJRqUV%2B0wXtjGSfg6ihhbOgowPAiAyhmQPmBhPBRcytqD0bjtnZeG0DORNxx0guWDmf8UEvQ%3D%3D)

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


Generated: 28 Nov 23 16:15 UTC