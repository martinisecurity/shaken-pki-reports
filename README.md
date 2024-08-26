# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://ecosystemcompliance.martinisecurity.com/#:~:text=Approved%20Certificate%20Authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is broken int two parts:
1. One generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main [Zlint](https://github.com/martinisecurity/zlint) distribution but can be found here.
2. One generated with a custom script that eumerates the known STIR/SHAKEN certificates and asses each repository against the current rule set . The source for this test can be found here while the report itself can be found [here](REPOS/README.md).

## Summary

### Leaf Certificates

- 13909 certificates were included in the corpus being tested
- 1503 certificates in the corpus were skipped because they are duplicates
- 10913 certificates in the corpus were skipped because they are expired
- 473 certificates in the corpus were skipped because they are not currently trusted
- 1020 certificates being tested against the remaining rules
- 1.74 issues on average found in unexpired, trusted, and non-compliant certificates
- 48.53% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 1.08% of certificates are too old to be assessed against currently enforced expectations
- 224 days is the average remaining validity for the certificates in the corpus
- 224 days is the average initial validity for the certificates in the corpus
- 399 certificates expire in the next 30 days
- 18.23 average number of unexpired certificates per OCN observed
- 763 unique OCNs observed in unexpired and valid certificate corpus

### CA Certificates

- 51 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 7 certificates in the corpus were skipped because they are not currently trusted
- 44 certificates being tested against the remaining rules
- 2.15 issues on average found in unexpired, trusted, and non-compliant certificates
- 45.45% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 65.91% of certificates are too old to be assessed against currently enforced expectations
- 5659 days is the average remaining validity for the certificates in the corpus
- 5585 days is the average initial validity for the certificates in the corpus
- 1 certificates expire in the next 30 days

## Certificate Repository URL

- 48.24% of certificate repository URLs contain one or more Error level issue
- 49.31% of certificates repository URLs contain one or more Warning level issue
- 0.00% of certificates repository URLs contain one or more Notice level issue

## Details

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

### Leaf Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#leaf-certificates) | 31 (3.04%) | 31 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#leaf-certificates) | 5 (0.49%) | 5 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#leaf-certificates) | 386 (37.84%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#leaf-certificates) | 64 (6.27%) | 64 (100.00%) | 0 (0.00%) | 0 (0.00%) | 11 (17.19%) |
| [NetNumber](CERTS/NetNumber/README.md#leaf-certificates) | 6 (0.59%) | 6 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Neustar](CERTS/Neustar/README.md#leaf-certificates) | 145 (14.22%) | 145 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Peeringhub](CERTS/Peeringhub/README.md#leaf-certificates) | 26 (2.55%) | 3 (11.54%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#leaf-certificates) | 12 (1.18%) | 12 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Sansay](CERTS/Sansay/README.md#leaf-certificates) | 215 (21.08%) | 215 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#leaf-certificates) | 2 (0.20%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Telonium](CERTS/Telonium/README.md#leaf-certificates) | 21 (2.06%) | 5 (23.81%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#leaf-certificates) | 107 (10.49%) | 7 (6.54%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| **Total** | 1020 (100.00%) | 495 (48.53%) | 0 (0.00%) | 0 (0.00%) | 11 (1.08%) |

### CA Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [CTIA](CERTS/CTIA/README.md#ca-certificates) | 1 (2.27%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Comcast](CERTS/Comcast/README.md#ca-certificates) | 2 (4.55%) | 1 (50.00%) | 0 (0.00%) | 0 (0.00%) | 1 (50.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#ca-certificates) | 3 (6.82%) | 2 (66.67%) | 0 (0.00%) | 0 (0.00%) | 2 (66.67%) |
| [Martini Security](CERTS/Martini_Security/README.md#ca-certificates) | 5 (11.36%) | 1 (20.00%) | 0 (0.00%) | 0 (0.00%) | 3 (60.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#ca-certificates) | 4 (9.09%) | 4 (100.00%) | 0 (0.00%) | 0 (0.00%) | 2 (50.00%) |
| [NetNumber](CERTS/NetNumber/README.md#ca-certificates) | 3 (6.82%) | 2 (66.67%) | 0 (0.00%) | 0 (0.00%) | 3 (100.00%) |
| [Neustar](CERTS/Neustar/README.md#ca-certificates) | 6 (13.64%) | 1 (16.67%) | 0 (0.00%) | 0 (0.00%) | 6 (100.00%) |
| [Peeringhub](CERTS/Peeringhub/README.md#ca-certificates) | 2 (4.55%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#ca-certificates) | 3 (6.82%) | 2 (66.67%) | 0 (0.00%) | 0 (0.00%) | 2 (66.67%) |
| [SOMOS](CERTS/SOMOS/README.md#ca-certificates) | 1 (2.27%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Sansay](CERTS/Sansay/README.md#ca-certificates) | 2 (4.55%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#ca-certificates) | 4 (9.09%) | 1 (25.00%) | 0 (0.00%) | 0 (0.00%) | 4 (100.00%) |
| [Telonium](CERTS/Telonium/README.md#ca-certificates) | 5 (11.36%) | 3 (60.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#ca-certificates) | 3 (6.82%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (66.67%) |
| **Total** | 44 (100.00%) | 20 (45.45%) | 0 (0.00%) | 0 (0.00%) | 29 (65.91%) |

### Key

| Type | Description |
|------|-------------|
| Errors | Tests in which the specifications are unambiguous on what the expected behavior must be. |
| Warnings | Tests in which the specifications are ambiguous or are provide only a recommendation. |
| Notices | Tests in which industry best practices are not followed. |
| Not Effective | Tests that exist in the current specifications but were not in effect at the time of issuance. |


Generated: 26 Aug 24 18:03 UTC