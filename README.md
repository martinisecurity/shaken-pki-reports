# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://ecosystemcompliance.martinisecurity.com/#:~:text=Approved%20Certificate%20Authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is broken int two parts:
1. One generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main [Zlint](https://github.com/martinisecurity/zlint) distribution but can be found here.
2. One generated with a custom script that eumerates the known STIR/SHAKEN certificates and asses each repository against the current rule set . The source for this test can be found here while the report itself can be found [here](REPOS/README.md).

## Summary

### Leaf Certificates

- 827 certificates were included in the corpus being tested
- 126 certificates in the corpus were skipped because they are duplicates
- 363 certificates in the corpus were skipped because they are expired
- 2 certificates in the corpus were skipped because they are not currently trusted
- 336 certificates being tested against the remaining rules
- 4.04 issues on average found in unexpired, trusted, and non-compliant certificates
- 93.75% of certificates contain one or more Error level issue
- 65.48% of certificates contain one or more Warning level issue
- 4.46% of certificates contain one or more Notice level issue
- 16.07% of certificates are too old to be assessed against currently enforced expectations
- 303 days is the average remaining validity for the certificates in the corpus
- 304 days is the average initial validity for the certificates in the corpus
- 148 certificates expire in the next 30 days

### CA Certificates

- 33 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 2 certificates in the corpus were skipped because they are expired
- 2 certificates in the corpus were skipped because they are not currently trusted
- 29 certificates being tested against the remaining rules
- 2.05 issues on average found in unexpired, trusted, and non-compliant certificates
- 55.17% of certificates contain one or more Error level issue
- 58.62% of certificates contain one or more Warning level issue
- 3.45% of certificates contain one or more Notice level issue
- 75.86% of certificates are too old to be assessed against currently enforced expectations
- 5741 days is the average remaining validity for the certificates in the corpus
- 5630 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

## Certificate Repository

- 51.49% of certificate repositories contain one or more Error level issue
- 77.68% of certificates repositories contain one or more Warning level issue
- 0.00% of certificates repositories contain one or more Notice level issue

## Details

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

### Leaf Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#leaf-certificates) | 35 (10.42%) | 35 (100.00%) | 35 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#leaf-certificates) | 1 (0.30%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#leaf-certificates) | 22 (6.55%) | 7 (31.82%) | 0 (0.00%) | 12 (54.55%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#leaf-certificates) | 22 (6.55%) | 22 (100.00%) | 1 (4.55%) | 0 (0.00%) | 21 (95.45%) |
| [NetNumber](CERTS/NetNumber/README.md#leaf-certificates) | 3 (0.89%) | 3 (100.00%) | 0 (0.00%) | 3 (100.00%) | 0 (0.00%) |
| [Neustar](CERTS/Neustar/README.md#leaf-certificates) | 96 (28.57%) | 90 (93.75%) | 32 (33.33%) | 0 (0.00%) | 32 (33.33%) |
| [Peeringhub](CERTS/Peeringhub/README.md#leaf-certificates) | 6 (1.79%) | 6 (100.00%) | 1 (16.67%) | 0 (0.00%) | 0 (0.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#leaf-certificates) | 3 (0.89%) | 3 (100.00%) | 3 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [Sansay](CERTS/Sansay/README.md#leaf-certificates) | 93 (27.68%) | 93 (100.00%) | 93 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#leaf-certificates) | 1 (0.30%) | 1 (100.00%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#leaf-certificates) | 54 (16.07%) | 54 (100.00%) | 54 (100.00%) | 0 (0.00%) | 1 (1.85%) |
| **Total** | 336 (100.00%) | 315 (93.75%) | 220 (65.48%) | 15 (4.46%) | 54 (16.07%) |

### CA Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#ca-certificates) | 2 (6.90%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#ca-certificates) | 2 (6.90%) | 1 (50.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#ca-certificates) | 3 (10.34%) | 0 (0.00%) | 0 (0.00%) | 1 (33.33%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#ca-certificates) | 2 (6.90%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [NetNumber](CERTS/NetNumber/README.md#ca-certificates) | 3 (10.34%) | 2 (66.67%) | 2 (66.67%) | 0 (0.00%) | 3 (100.00%) |
| [Neustar](CERTS/Neustar/README.md#ca-certificates) | 6 (20.69%) | 4 (66.67%) | 6 (100.00%) | 0 (0.00%) | 4 (66.67%) |
| [Peeringhub](CERTS/Peeringhub/README.md#ca-certificates) | 2 (6.90%) | 2 (100.00%) | 2 (100.00%) | 0 (0.00%) | 1 (50.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#ca-certificates) | 2 (6.90%) | 2 (100.00%) | 2 (100.00%) | 0 (0.00%) | 2 (100.00%) |
| [Sansay](CERTS/Sansay/README.md#ca-certificates) | 2 (6.90%) | 0 (0.00%) | 2 (100.00%) | 0 (0.00%) | 1 (50.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#ca-certificates) | 2 (6.90%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [TransNexus](CERTS/TransNexus/README.md#ca-certificates) | 3 (10.34%) | 3 (100.00%) | 3 (100.00%) | 0 (0.00%) | 3 (100.00%) |
| **Total** | 29 (100.00%) | 16 (55.17%) | 17 (58.62%) | 1 (3.45%) | 22 (75.86%) |

### Key

| Type | Description |
|------|-------------|
| Errors | Tests in which the specifications are unambiguous on what the expected behavior must be. |
| Warnings | Tests in which the specifications are ambiguous or are provide only a recommendation. |
| Notices | Tests in which industry best practices are not followed. |
| Not Effective | Tests that exist in the current specifications but were not in effect at the time of issuance. |


Generated: 06 Nov 22 17:28 UTC