# 12 Deployment and maintenance

## 12.1 Logs

- `logrus`: "github.com/Sirupsen/logrus"
- `seelog`: "github.com/cihub/seelog"

## 12.2 Errors and crashes

## 12.3 Deployment

### Daemons

### Supervisord

## 12.4 Backup and recovery

### Application Backup
#### Rsync installation + configuration

### MySQL backup

Hot backups: `master/slave`
Cold backups: `shell` script + rsync -> local server


### MySQL recovery
`mysql -u username -p database < backup.sql`


### Redis backup

Hot: `master/slave`
Cold: routinely saves cached data in memory to the database file on-disk.

### Redis recovery

Hot: like MySQL
Cold: copy redis file to redis root then `$ redis-server` 
