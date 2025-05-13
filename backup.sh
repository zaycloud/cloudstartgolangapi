#!/bin/bash

# Sätt datum för backup-filen
DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE="backup_${DATE}.sql"

# Gör backup av MySQL-databasen
mysqldump -h $MYSQL_HOST -u $MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE > $BACKUP_FILE

# Komprimera backup-filen
gzip $BACKUP_FILE

# Ladda upp till Backblaze S3
aws s3 cp ${BACKUP_FILE}.gz s3://$S3_BUCKET/backups/ --endpoint-url $S3_ENDPOINT

# Ta bort lokala filer
rm ${BACKUP_FILE}.gz 