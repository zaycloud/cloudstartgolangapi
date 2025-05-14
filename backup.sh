#!/bin/sh
set -e

echo "Starting backup of database $MYSQL_DATABASE..."

# Create backup filename with timestamp
BACKUP_FILE="/backup/backup-$(date +%Y%m%d-%H%M%S).sql"

# Perform MySQL backup
mysqldump -h $MYSQL_HOST -u $MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE > $BACKUP_FILE

# Compress the backup file
gzip $BACKUP_FILE
BACKUP_FILE="${BACKUP_FILE}.gz"

echo "Configuring mc..."
mc alias set b2 "$S3_ENDPOINT" "$AWS_ACCESS_KEY_ID" "$AWS_SECRET_ACCESS_KEY"

echo "Uploading to Backblaze S3..."
mc cp "$BACKUP_FILE" "b2/$S3_BUCKET/$(basename $BACKUP_FILE)"

# Clean up local files
rm $BACKUP_FILE

echo "Backup process completed successfully" 