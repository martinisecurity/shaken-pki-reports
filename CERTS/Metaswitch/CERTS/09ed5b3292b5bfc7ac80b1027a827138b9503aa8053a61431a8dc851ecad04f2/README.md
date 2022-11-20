# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Cspire SHAKEN Cert 6581

Tested At: 20 Nov 22 22:57 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 490 day(s)\
Subject: CN=Cspire SHAKEN Cert 6581, O=Cspire, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/6c54d88f5669c28848983d4572f11c0b896342ac

[View certificate details](https://understandingwebpki.com/?cert=MIICRjCCAe2gAwIBAgIQE3lY2DnuwU0MPjNdy%2BIj3TAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDMyNDIzNTI0NFoXDTI0MDMyMzIzNTI0NFowQDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBkNzcGlyZTEgMB4GA1UEAwwXQ3NwaXJlIFNIQUtFTiBDZXJ0IDY1ODEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS1it%2Ffie%2BxEYwKEiByYLAolMSwz53x2oBUB9DOdK%2FIZqPk82h%2B8Ou1nvO%2FZlSmnsG5C81gqI4CHdX2w7nVlTtqo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ2NTgxMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFGjSBWqoiWThMqlMxnbXceiQcx2XMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIGuhzlPpon4tSFuUYrsJRqUV%2B0wXtjGSfg6ihhbOgowPAiAyhmQPmBhPBRcytqD0bjtnZeG0DORNxx0guWDmf8UEvQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 20 Nov 22 22:57 UTC