import React, { useEffect, useState } from 'react';
import {
  Button,
  Form,
  Header,
  Label,
  Pagination,
  Segment,
  Select,
  Table,
  Popup,
  Icon,
  Dropdown,
} from 'semantic-ui-react';
import {
  API,
  copy,
  isAdmin,
  showError,
  showSuccess,
  showWarning,
  timestamp2string,
} from '../helpers';
import { useTranslation } from 'react-i18next';

import { ITEMS_PER_PAGE } from '../constants';
import { renderColorLabel, renderQuota } from '../helpers/render';
import { Link } from 'react-router-dom';

function renderTimestamp(timestamp, request_id) {
  return (
    <code
      onClick={async () => {
        if (await copy(request_id)) {
          showSuccess(`Request ID copied: ${request_id}`);
        } else {
          showWarning(`Failed to copy request ID: ${request_id}`);
        }
      }}
      style={{ cursor: 'pointer' }}
    >
      {timestamp2string(timestamp)}
    </code>
  );
}

const MODE_OPTIONS = [
  { key: 'all', text: 'All Users', value: 'all' },
  { key: 'self', text: 'Current User', value: 'self' },
];

function renderType(type) {
  switch (type) {
    case 1:
      return (
        <Label basic color='green'>
          Recharge
        </Label>
      );
    case 2:
      return (
        <Label basic color='olive'>
          Consumed
        </Label>
      );
    case 3:
      return (
        <Label basic color='orange'>
          Management
        </Label>
      );
    case 4:
      return (
        <Label basic color='purple'>
          System
        </Label>
      );
    case 5:
      return (
        <Label basic color='violet'>
          Test
        </Label>
      );
    default:
      return (
        <Label basic color='black'>
          Unknown
        </Label>
      );
  }
}



function renderDetail(log) {
  return (
    <>
      {log.content}
      <br />
      {log.is_stream && (
        <>
          <Label size={'mini'} color='pink'>
            Stream
          </Label>
        </>
      )}
      {log.system_prompt_reset && (
        <>
          <Label basic size={'mini'} color='red'>
            System Prompt Reset
          </Label>
        </>
      )}
    </>
  );
}

function renderLatency(elapsedTime) {
  if (!elapsedTime) return '';
  return `${elapsedTime} ms`;
}

function getSortIcon(columnKey, currentSortBy, currentSortOrder) {
  if (columnKey !== currentSortBy) {
    return null;
  }
  return currentSortOrder === 'asc' ? ' ↑' : ' ↓';
}

const LogsTable = () => {
  const { t } = useTranslation();
  const [logs, setLogs] = useState([]);
  const [showStat, setShowStat] = useState(false);
  const [loading, setLoading] = useState(true);
  const [activePage, setActivePage] = useState(1);
  const [searchKeyword, setSearchKeyword] = useState('');

  const [logType, setLogType] = useState(0);
  const isAdminUser = isAdmin();
  let now = new Date();
  let sevenDaysAgo = new Date(now.getTime() - 7 * 24 * 60 * 60 * 1000);
  const [inputs, setInputs] = useState({
    username: '',
    token_name: '',
    model_name: '',
    start_timestamp: timestamp2string(sevenDaysAgo.getTime() / 1000),
    end_timestamp: timestamp2string(now.getTime() / 1000 + 3600),
    channel: '',
  });
  const {
    username,
    token_name,
    model_name,
    start_timestamp,
    end_timestamp,
    channel,
  } = inputs;

  const [stat, setStat] = useState({
    quota: 0,
    token: 0,
  });
  const [isStatRefreshing, setIsStatRefreshing] = useState(false);
  const [userOptions, setUserOptions] = useState([]);
  const [userSearchLoading, setUserSearchLoading] = useState(false);
  const [sortBy, setSortBy] = useState('');
  const [sortOrder, setSortOrder] = useState('desc');
  const [sortLoading, setSortLoading] = useState(false);

  const LOG_OPTIONS = [
    { key: '0', text: t('log.type.all'), value: 0 },
    { key: '1', text: t('log.type.topup'), value: 1 },
    { key: '2', text: t('log.type.usage'), value: 2 },
    { key: '3', text: t('log.type.admin'), value: 3 },
    { key: '4', text: t('log.type.system'), value: 4 },
    { key: '5', text: t('log.type.test'), value: 5 },
  ];

  const handleInputChange = (_, { name, value }) => {
    setInputs((inputs) => ({ ...inputs, [name]: value }));
  };

  const searchUsers = async (searchQuery) => {
    if (!searchQuery.trim()) {
      setUserOptions([]);
      return;
    }

    setUserSearchLoading(true);
    try {
      const res = await API.get(`/api/user/search?keyword=${searchQuery}`);
      const { success, data } = res.data;
      if (success) {
        const options = data.map(user => ({
          key: user.id,
          value: user.username,
          text: `${user.display_name || user.username} (@${user.username})`,
          content: (
            <div style={{ display: 'flex', flexDirection: 'column' }}>
              <div style={{ fontWeight: 'bold' }}>
                {user.display_name || user.username}
              </div>
              <div style={{ fontSize: '0.9em', color: '#666' }}>
                @{user.username} • ID: {user.id}
              </div>
            </div>
          )
        }));
        setUserOptions(options);
      }
    } catch (error) {
      console.error('Failed to search users:', error);
    } finally {
      setUserSearchLoading(false);
    }
  };

  const handleStatRefresh = async () => {
    setIsStatRefreshing(true);
    try {
      if (isAdminUser) {
        await getLogStat();
      } else {
        await getLogSelfStat();
      }
    } finally {
      setIsStatRefreshing(false);
    }
  };

  const getLogSelfStat = async () => {
    let localStartTimestamp = Date.parse(start_timestamp) / 1000;
    let localEndTimestamp = Date.parse(end_timestamp) / 1000;
    let res = await API.get(
      `/api/log/self/stat?type=${logType}&token_name=${token_name}&model_name=${model_name}&start_timestamp=${localStartTimestamp}&end_timestamp=${localEndTimestamp}`
    );
    const { success, message, data } = res.data;
    if (success) {
      setStat(data);
    } else {
      showError(message);
    }
  };

  const getLogStat = async () => {
    let localStartTimestamp = Date.parse(start_timestamp) / 1000;
    let localEndTimestamp = Date.parse(end_timestamp) / 1000;
    let res = await API.get(
      `/api/log/stat?type=${logType}&username=${username}&token_name=${token_name}&model_name=${model_name}&start_timestamp=${localStartTimestamp}&end_timestamp=${localEndTimestamp}&channel=${channel}`
    );
    const { success, message, data } = res.data;
    if (success) {
      setStat(data);
    } else {
      showError(message);
    }
  };

  const handleEyeClick = async () => {
    if (!showStat) {
      if (isAdminUser) {
        await getLogStat();
      } else {
        await getLogSelfStat();
      }
    }
    setShowStat(!showStat);
  };

  const showUserTokenQuota = () => {
    return logType !== 5;
  };

  const loadLogs = async (startIdx) => {
    let url = '';
    let localStartTimestamp = Date.parse(start_timestamp) / 1000;
    let localEndTimestamp = Date.parse(end_timestamp) / 1000;
    let sortParams = '';
    if (sortBy) {
      sortParams = `&sort_by=${sortBy}&sort_order=${sortOrder}`;
    }
    if (isAdminUser) {
      url = `/api/log/?p=${startIdx}&type=${logType}&username=${username}&token_name=${token_name}&model_name=${model_name}&start_timestamp=${localStartTimestamp}&end_timestamp=${localEndTimestamp}&channel=${channel}${sortParams}`;
    } else {
      url = `/api/log/self?p=${startIdx}&type=${logType}&token_name=${token_name}&model_name=${model_name}&start_timestamp=${localStartTimestamp}&end_timestamp=${localEndTimestamp}${sortParams}`;
    }
    const res = await API.get(url);
    const { success, message, data } = res.data;
    if (success) {
      if (startIdx === 0) {
        setLogs(data);
      } else {
        let newLogs = [...logs];
        newLogs.splice(startIdx * ITEMS_PER_PAGE, data.length, ...data);
        setLogs(newLogs);
      }
    } else {
      showError(message);
    }
    setLoading(false);
  };

  const onPaginationChange = (_, { activePage }) => {
    (async () => {
      if (activePage === Math.ceil(logs.length / ITEMS_PER_PAGE) + 1) {
        // In this case we have to load more data and then append them.
        await loadLogs(activePage - 1);
      }
      setActivePage(activePage);
    })();
  };

  const refresh = async () => {
    setLoading(true);
    setActivePage(1);
    await loadLogs(0);
  };

  useEffect(() => {
    refresh();
  }, [logType, sortBy, sortOrder]);



  const sortLog = async (key) => {
    // Prevent multiple sort requests
    if (sortLoading) return;

    // Toggle sort order if clicking the same column
    let newSortOrder = 'desc';
    if (sortBy === key && sortOrder === 'desc') {
      newSortOrder = 'asc';
    }

    setSortBy(key);
    setSortOrder(newSortOrder);
    setActivePage(1);
    setSortLoading(true);

    try {
      // Reload data with new sorting
      await loadLogs(0);
    } finally {
      setSortLoading(false);
    }
  };

  return (
    <>
      <Header as='h3'>
        {t('log.usage_details')}（{t('log.total_quota')}：
        {showStat && (
          <>
            {renderQuota(stat.quota, t)}
            <Button
              size='mini'
              circular
              icon='refresh'
              onClick={handleStatRefresh}
              loading={isStatRefreshing}
              disabled={isStatRefreshing}
              style={{
                marginLeft: '8px',
                padding: '4px',
                minHeight: '20px',
                minWidth: '20px',
                fontSize: '10px'
              }}
              title={t('log.refresh_quota', 'Refresh quota data')}
            />
          </>
        )}
        {!showStat && (
          <span
            onClick={handleEyeClick}
            style={{ cursor: 'pointer', color: 'gray' }}
          >
            {t('log.click_to_view')}
          </span>
        )}
        ）
      </Header>
      <Form>
        <Form.Group>
          <Form.Input
            fluid
            label={t('log.table.token_name')}
            size={'small'}
            width={3}
            value={token_name}
            placeholder={t('log.table.token_name_placeholder')}
            name='token_name'
            onChange={handleInputChange}
          />
          <Form.Input
            fluid
            label={t('log.table.model_name')}
            size={'small'}
            width={3}
            value={model_name}
            placeholder={t('log.table.model_name_placeholder')}
            name='model_name'
            onChange={handleInputChange}
          />
          <Form.Input
            fluid
            label={t('log.table.start_time')}
            size={'small'}
            width={4}
            value={start_timestamp}
            type='datetime-local'
            name='start_timestamp'
            onChange={handleInputChange}
          />
          <Form.Input
            fluid
            label={t('log.table.end_time')}
            size={'small'}
            width={4}
            value={end_timestamp}
            type='datetime-local'
            name='end_timestamp'
            onChange={handleInputChange}
          />
          <Form.Button
            fluid
            label={t('log.buttons.query')}
            size={'small'}
            width={2}
            onClick={refresh}
          >
            {t('log.buttons.submit')}
          </Form.Button>
        </Form.Group>
        {isAdminUser && (
          <>
            <Form.Group>
              <Form.Input
                fluid
                label={t('log.table.channel_id')}
                size={'small'}
                width={3}
                value={channel}
                placeholder={t('log.table.channel_id_placeholder')}
                name='channel'
                onChange={handleInputChange}
              />
              <Form.Field width={3}>
                <label>{t('log.table.username')}</label>
                <Dropdown
                  fluid
                  selection
                  search
                  clearable
                  allowAdditions
                  value={username}
                  placeholder={t('log.table.username_placeholder')}
                  options={userOptions}
                  onSearchChange={(_, { searchQuery }) => searchUsers(searchQuery)}
                  onChange={(_, { value }) => handleInputChange(_, { name: 'username', value })}
                  loading={userSearchLoading}
                  noResultsMessage={t('log.no_users_found', 'No users found')}
                  additionLabel={t('log.use_username', 'Use username: ')}
                  onAddItem={(_, { value }) => {
                    const newOption = {
                      key: value,
                      value: value,
                      text: value
                    };
                    setUserOptions([...userOptions, newOption]);
                  }}
                />
              </Form.Field>
            </Form.Group>
          </>
        )}
        <Form.Input
          icon='search'
          placeholder={t('log.search')}
          value={searchKeyword}
          onChange={(_, { value }) => setSearchKeyword(value)}
        />
      </Form>
      <Table basic={'very'} compact size='small'>
        <Table.Header>
          <Table.Row>
            <Table.HeaderCell
              style={{ cursor: 'pointer' }}
              onClick={() => {
                sortLog('created_time');
              }}
              width={3}
            >
              {t('log.table.time')}
            </Table.HeaderCell>
            {isAdminUser && (
              <Table.HeaderCell
                style={{ cursor: 'pointer' }}
                onClick={() => {
                  sortLog('channel');
                }}
                width={1}
              >
                {t('log.table.channel')}
              </Table.HeaderCell>
            )}
            <Table.HeaderCell
              style={{ cursor: 'pointer' }}
              onClick={() => {
                sortLog('type');
              }}
              width={1}
            >
              {t('log.table.type')}
            </Table.HeaderCell>
            <Table.HeaderCell
              style={{ cursor: 'pointer' }}
              onClick={() => {
                sortLog('model_name');
              }}
              width={2}
            >
              {t('log.table.model')}
            </Table.HeaderCell>
            {showUserTokenQuota() && (
              <>
                {isAdminUser && (
                  <Table.HeaderCell
                    style={{ cursor: 'pointer' }}
                    onClick={() => {
                      sortLog('username');
                    }}
                    width={2}
                  >
                    {t('log.table.username')}
                  </Table.HeaderCell>
                )}
                <Table.HeaderCell
                  style={{ cursor: 'pointer' }}
                  onClick={() => {
                    sortLog('token_name');
                  }}
                  width={2}
                >
                  {t('log.table.token_name')}
                </Table.HeaderCell>
                <Table.HeaderCell
                  style={{ cursor: sortLoading ? 'wait' : 'pointer', opacity: sortLoading ? 0.6 : 1 }}
                  onClick={() => {
                    sortLog('prompt_tokens');
                  }}
                  width={1}
                >
                  {t('log.table.prompt_tokens')}{getSortIcon('prompt_tokens', sortBy, sortOrder)}
                  {sortLoading && sortBy === 'prompt_tokens' && <span> ⏳</span>}
                </Table.HeaderCell>
                <Table.HeaderCell
                  style={{ cursor: sortLoading ? 'wait' : 'pointer', opacity: sortLoading ? 0.6 : 1 }}
                  onClick={() => {
                    sortLog('completion_tokens');
                  }}
                  width={1}
                >
                  {t('log.table.completion_tokens')}{getSortIcon('completion_tokens', sortBy, sortOrder)}
                  {sortLoading && sortBy === 'completion_tokens' && <span> ⏳</span>}
                </Table.HeaderCell>
                <Table.HeaderCell
                  style={{ cursor: sortLoading ? 'wait' : 'pointer', opacity: sortLoading ? 0.6 : 1 }}
                  onClick={() => {
                    sortLog('quota');
                  }}
                  width={1}
                >
                  {t('log.table.quota')}{getSortIcon('quota', sortBy, sortOrder)}
                  {sortLoading && sortBy === 'quota' && <span> ⏳</span>}
                </Table.HeaderCell>
                <Table.HeaderCell
                  style={{ cursor: sortLoading ? 'wait' : 'pointer', opacity: sortLoading ? 0.6 : 1 }}
                  onClick={() => {
                    sortLog('elapsed_time');
                  }}
                  width={1}
                >
                  {t('log.table.latency')}{getSortIcon('elapsed_time', sortBy, sortOrder)}
                  {sortLoading && sortBy === 'elapsed_time' && <span> ⏳</span>}
                </Table.HeaderCell>
              </>
            )}
            <Table.HeaderCell>{t('log.table.detail')}</Table.HeaderCell>
          </Table.Row>
        </Table.Header>

        <Table.Body>
          {logs
            .slice(
              (activePage - 1) * ITEMS_PER_PAGE,
              activePage * ITEMS_PER_PAGE
            )
            .map((log) => {
              if (log.deleted) return <></>;
              return (
                <Table.Row key={log.id}>
                  <Table.Cell>
                    {renderTimestamp(log.created_at, log.request_id)}
                  </Table.Cell>
                  {isAdminUser && (
                    <Table.Cell>
                      {log.channel ? (
                        <Label
                          basic
                          as={Link}
                          to={`/channel/edit/${log.channel}`}
                        >
                          {log.channel}
                        </Label>
                      ) : (
                        ''
                      )}
                    </Table.Cell>
                  )}
                  <Table.Cell>{renderType(log.type)}</Table.Cell>
                  <Table.Cell>
                    {log.model_name ? renderColorLabel(log.model_name) : ''}
                  </Table.Cell>
                  {showUserTokenQuota() && (
                    <>
                      {isAdminUser && (
                        <Table.Cell>
                          {log.username ? (
                            <Label
                              basic
                              as={Link}
                              to={`/user/edit/${log.user_id}`}
                            >
                              {log.username}
                            </Label>
                          ) : (
                            ''
                          )}
                        </Table.Cell>
                      )}
                      <Table.Cell>
                        {log.token_name ? renderColorLabel(log.token_name) : ''}
                      </Table.Cell>

                      <Table.Cell>
                        {log.prompt_tokens ? log.prompt_tokens : ''}
                      </Table.Cell>
                      <Table.Cell>
                        {log.completion_tokens ? log.completion_tokens : ''}
                      </Table.Cell>
                      <Table.Cell>
                        {log.quota ? renderQuota(log.quota, t, 6) : 'free'}
                      </Table.Cell>
                      <Table.Cell>
                        {renderLatency(log.elapsed_time)}
                      </Table.Cell>
                    </>
                  )}

                  <Table.Cell>{renderDetail(log)}</Table.Cell>
                </Table.Row>
              );
            })}
        </Table.Body>

        <Table.Footer>
          <Table.Row>
            <Table.HeaderCell colSpan={'10'}>
              <Select
                placeholder={t('log.type.select')}
                options={LOG_OPTIONS}
                style={{ marginRight: '8px' }}
                name='logType'
                value={logType}
                onChange={(_, { value }) => {
                  setLogType(value);
                }}
              />
              <Button size='small' onClick={refresh} loading={loading}>
                {t('log.buttons.refresh')}
              </Button>
              <Pagination
                floated='right'
                activePage={activePage}
                onPageChange={onPaginationChange}
                size='small'
                siblingRange={1}
                totalPages={
                  Math.ceil(logs.length / ITEMS_PER_PAGE) +
                  (logs.length % ITEMS_PER_PAGE === 0 ? 1 : 0)
                }
              />
            </Table.HeaderCell>
          </Table.Row>
        </Table.Footer>
      </Table>
    </>
  );
};

export default LogsTable;
