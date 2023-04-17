CREATE TRIGGER ATT_JSON_To_Table
ON dbo.ATTMessagingJSON
AFTER INSERT
NOT FOR REPLICATION
AS
BEGIN
    DECLARE
        @json nvarchar(max)
        SET NOCOUNT ON
        SELECT @json = INSERTED.[payload] FROM INSERTED
    BEGIN
        INSERT INTO dbo.ATTMessaging (subscriberId, msisdn, ban, messageId, messageContext, messageTime, "from", "to", groupMessage, direction, "subject", contentType, textContent, "textSize", attachment)
        SELECT subscriberId, msisdn, ban, messageId, messageContext, messageTime, "from", "to", groupMessage, direction, "subject", contentType, textContent, "textSize", attachment
        FROM OPENJSON (@json)
        WITH (subscriberId varchar(100), msisdn varchar(50), ban varchar(50), messageId varchar(100), messageContext varchar(50), messageTime datetime, "from" nvarchar(100), "to" nvarchar(max) AS JSON, groupMessage varchar(20), direction varchar(10), "subject" nvarchar(max), contentType varchar(50), textContent nvarchar(max), "textSize" varchar(100), attachment nvarchar(max) AS JSON)
    END
END
