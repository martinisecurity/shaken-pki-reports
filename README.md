# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://ecosystemcompliance.martinisecurity.com/#:~:text=Approved%20Certificate%20Authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is broken int two parts:
1. One generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main [Zlint](https://github.com/martinisecurity/zlint) distribution but can be found here.
2. One generated with a custom script that eumerates the known STIR/SHAKEN certificates and asses each repository against the current rule set . The source for this test can be found here while the report itself can be found [here](REPOS/README.md).

## Summary

### Leaf Certificates

- 910 certificates were included in the corpus being tested
- 133 certificates in the corpus were skipped because they are duplicates
- 473 certificates in the corpus were skipped because they are expired
- 2 certificates in the corpus were skipped because they are not currently trusted
- 302 certificates being tested against the remaining rules
- 4.19 issues on average found in unexpired, trusted, and non-compliant certificates
- 90.40% of certificates contain one or more Error level issue
- 57.62% of certificates contain one or more Warning level issue
- 2.98% of certificates contain one or more Notice level issue
- 17.55% of certificates are too old to be assessed against currently enforced expectations
- 339 days is the average remaining validity for the certificates in the corpus
- 340 days is the average initial validity for the certificates in the corpus
- 108 certificates expire in the next 30 days

### CA Certificates

- 34 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 1 certificates in the corpus were skipped because they are expired
- 2 certificates in the corpus were skipped because they are not currently trusted
- 31 certificates being tested against the remaining rules
- 2.05 issues on average found in unexpired, trusted, and non-compliant certificates
- 51.61% of certificates contain one or more Error level issue
- 54.84% of certificates contain one or more Warning level issue
- 3.23% of certificates contain one or more Notice level issue
- 74.19% of certificates are too old to be assessed against currently enforced expectations
- 5865 days is the average remaining validity for the certificates in the corpus
- 5797 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

## Certificate Repository URL

- 55.96% of certificate repository URLs contain one or more Error level issue
- 83.77% of certificates repository URLs contain one or more Warning level issue
- 0.00% of certificates repository URLs contain one or more Notice level issue

## Details

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

### Leaf Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#leaf-certificates) | 25 (8.28%) | 25 (100.00%) | 25 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#leaf-certificates) | 2 (0.66%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#leaf-certificates) | 25 (8.28%) | 0 (0.00%) | 0 (0.00%) | 5 (20.00%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#leaf-certificates) | 23 (7.62%) | 23 (100.00%) | 1 (4.35%) | 0 (0.00%) | 22 (95.65%) |
| [NetNumber](CERTS/NetNumber/README.md#leaf-certificates) | 4 (1.32%) | 4 (100.00%) | 0 (0.00%) | 4 (100.00%) | 0 (0.00%) |
| [Neustar](CERTS/Neustar/README.md#leaf-certificates) | 100 (33.11%) | 96 (96.00%) | 30 (30.00%) | 0 (0.00%) | 30 (30.00%) |
| [Peeringhub](CERTS/Peeringhub/README.md#leaf-certificates) | 6 (1.99%) | 6 (100.00%) | 1 (16.67%) | 0 (0.00%) | 0 (0.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#leaf-certificates) | 3 (0.99%) | 3 (100.00%) | 3 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [Sansay](CERTS/Sansay/README.md#leaf-certificates) | 89 (29.47%) | 89 (100.00%) | 89 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#leaf-certificates) | 1 (0.33%) | 1 (100.00%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#leaf-certificates) | 24 (7.95%) | 24 (100.00%) | 24 (100.00%) | 0 (0.00%) | 1 (4.17%) |
| **Total** | 302 (100.00%) | 273 (90.40%) | 174 (57.62%) | 9 (2.98%) | 53 (17.55%) |

### CA Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#ca-certificates) | 2 (6.45%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#ca-certificates) | 2 (6.45%) | 1 (50.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#ca-certificates) | 3 (9.68%) | 0 (0.00%) | 0 (0.00%) | 1 (33.33%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#ca-certificates) | 2 (6.45%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [NetNumber](CERTS/NetNumber/README.md#ca-certificates) | 3 (9.68%) | 2 (66.67%) | 2 (66.67%) | 0 (0.00%) | 3 (100.00%) |
| [Neustar](CERTS/Neustar/README.md#ca-certificates) | 6 (19.35%) | 4 (66.67%) | 6 (100.00%) | 0 (0.00%) | 4 (66.67%) |
| [Peeringhub](CERTS/Peeringhub/README.md#ca-certificates) | 2 (6.45%) | 2 (100.00%) | 2 (100.00%) | 0 (0.00%) | 1 (50.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#ca-certificates) | 2 (6.45%) | 2 (100.00%) | 2 (100.00%) | 0 (0.00%) | 2 (100.00%) |
| [Sansay](CERTS/Sansay/README.md#ca-certificates) | 2 (6.45%) | 0 (0.00%) | 2 (100.00%) | 0 (0.00%) | 1 (50.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#ca-certificates) | 3 (9.68%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 3 (100.00%) |
| [TransNexus](CERTS/TransNexus/README.md#ca-certificates) | 4 (12.90%) | 3 (75.00%) | 3 (75.00%) | 0 (0.00%) | 3 (75.00%) |
| **Total** | 31 (100.00%) | 16 (51.61%) | 17 (54.84%) | 1 (3.23%) | 23 (74.19%) |

### Key

| Type | Description |
|------|-------------|
| Errors | Tests in which the specifications are unambiguous on what the expected behavior must be. |
| Warnings | Tests in which the specifications are ambiguous or are provide only a recommendation. |
| Notices | Tests in which industry best practices are not followed. |
| Not Effective | Tests that exist in the current specifications but were not in effect at the time of issuance. |


Generated: 17 Nov 22 19:20 UTC