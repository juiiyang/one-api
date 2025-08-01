import React from 'react';
import {enqueueSnackbar} from 'notistack';
import {snackbarConstants} from 'constants/SnackbarConstants';
import {API} from './api';

export function getSystemName() {
    let system_name = localStorage.getItem('system_name');
    if (!system_name) return 'One API';
    return system_name;
}

export function isMobile() {
    return window.innerWidth <= 600;
}

// eslint-disable-next-line
export function SnackbarHTMLContent({htmlContent}) {
    return <div dangerouslySetInnerHTML={{__html: htmlContent}}/>;
}

export function getSnackbarOptions(variant) {
    let options = snackbarConstants.Common[variant];
    if (isMobile()) {
        // 合并 options 和 snackbarConstants.Mobile
        options = {...options, ...snackbarConstants.Mobile};
    }
    return options;
}

export function showError(error) {
    if (error.message) {
        if (error.name === 'AxiosError') {
            switch (error.response.status) {
                case 429:
                    enqueueSnackbar('错误：请求次数过多，请稍后再试！', getSnackbarOptions('ERROR'));
                    break;
                case 500:
                    enqueueSnackbar('错误：服务器内部错误，请联系管理员！', getSnackbarOptions('ERROR'));
                    break;
                case 405:
                    enqueueSnackbar('本站仅作演示之用，无服务端！', getSnackbarOptions('INFO'));
                    break;
                default:
                    enqueueSnackbar('错误：' + error.message, getSnackbarOptions('ERROR'));
            }
            return;
        }
    } else {
        enqueueSnackbar('错误：' + error, getSnackbarOptions('ERROR'));
    }
}

export function showNotice(message, isHTML = false) {
    if (isHTML) {
        enqueueSnackbar(<SnackbarHTMLContent htmlContent={message}/>, getSnackbarOptions('NOTICE'));
    } else {
        enqueueSnackbar(message, getSnackbarOptions('NOTICE'));
    }
}

export function showWarning(message) {
    enqueueSnackbar(message, getSnackbarOptions('WARNING'));
}

export function showSuccess(message) {
    enqueueSnackbar(message, getSnackbarOptions('SUCCESS'));
}

export function showInfo(message) {
    enqueueSnackbar(message, getSnackbarOptions('INFO'));
}

export async function getOAuthState() {
    const res = await API.get('/api/oauth/state');
    const {success, message, data} = res.data;
    if (success) {
        return data;
    } else {
        showError(message);
        return '';
    }
}

export async function onGitHubOAuthClicked(github_client_id, openInNewTab = false) {
    const state = await getOAuthState();
    if (!state) return;
    let url = `https://github.com/login/oauth/authorize?client_id=${github_client_id}&state=${state}&scope=user:email`;
    if (openInNewTab) {
        window.open(url);
    } else {
        window.location.href = url;
    }
}

export async function onLarkOAuthClicked(lark_client_id) {
    const state = await getOAuthState();
    if (!state) return;
    let redirect_uri = `${window.location.origin}/oauth/lark`;
    window.open(`https://accounts.feishu.cn/open-apis/authen/v1/authorize?redirect_uri=${redirect_uri}&client_id=${lark_client_id}&state=${state}`);
}

export async function onOidcClicked(auth_url, client_id, openInNewTab = false) {
    const state = await getOAuthState();
    if (!state) return;
    const redirect_uri = `${window.location.origin}/oauth/oidc`;
    const response_type = "code";
    const scope = "openid profile email";
    const url = `${auth_url}?client_id=${client_id}&redirect_uri=${redirect_uri}&response_type=${response_type}&scope=${scope}&state=${state}`;
    if (openInNewTab) {
        window.open(url);
    } else {
        window.location.href = url;
    }
}

export function isAdmin() {
    let user = localStorage.getItem('user');
    if (!user) return false;
    user = JSON.parse(user);
    return user.role >= 10;
}

export function isRoot() {
    let user = localStorage.getItem('user');
    if (!user) return false;
    user = JSON.parse(user);
    return user.role >= 100;
}

export function timestamp2string(timestamp) {
    let date = new Date(timestamp * 1000);
    let year = date.getFullYear().toString();
    let month = (date.getMonth() + 1).toString();
    let day = date.getDate().toString();
    let hour = date.getHours().toString();
    let minute = date.getMinutes().toString();
    let second = date.getSeconds().toString();
    if (month.length === 1) {
        month = '0' + month;
    }
    if (day.length === 1) {
        day = '0' + day;
    }
    if (hour.length === 1) {
        hour = '0' + hour;
    }
    if (minute.length === 1) {
        minute = '0' + minute;
    }
    if (second.length === 1) {
        second = '0' + second;
    }
    return year + '-' + month + '-' + day + ' ' + hour + ':' + minute + ':' + second;
}

export function calculateQuota(quota = 0, digits = 2) {
    let quotaPerUnit = localStorage.getItem('quota_per_unit');
    quotaPerUnit = parseFloat(quotaPerUnit || '500000');

    return (quota / quotaPerUnit).toFixed(digits);
}

export function renderQuota(quota, digits = 2) {
    let displayInCurrency = localStorage.getItem('display_in_currency');
    displayInCurrency = displayInCurrency === 'true';
    if (displayInCurrency) {
        return '$' + calculateQuota(quota, digits);
    }
    return renderNumber(quota);
}

export const verifyJSON = (str) => {
    try {
        JSON.parse(str);
    } catch (e) {
        return false;
    }
    return true;
};

export function renderNumber(num) {
    if (num >= 1000000000) {
        return (num / 1000000000).toFixed(1) + 'B';
    } else if (num >= 1000000) {
        return (num / 1000000).toFixed(1) + 'M';
    } else if (num >= 10000) {
        return (num / 1000).toFixed(1) + 'k';
    } else {
        return num;
    }
}

/**
 * Renders a number with tooltip showing exact value when abbreviated.
 * WARNING: Returns JSX element for abbreviated numbers, plain value otherwise.
 * Only use in React components, not for plain text contexts.
 * @param {number} num - The number to render
 * @returns {JSX.Element|string|number} JSX span with tooltip for large numbers, plain value otherwise
 */
export function renderNumberWithTooltip(num) {
    const abbreviated = renderNumber(num);
    const exact = num.toLocaleString();

    // If the number was abbreviated, return JSX with tooltip
    if (num >= 10000) {
        return (
            <span title={exact} style={{ cursor: 'help', textDecoration: 'underline dotted' }}>
                {abbreviated}
            </span>
        );
    }

    // If not abbreviated, just return the number
    return abbreviated;
}

export function renderNumberForChart(num) {
    const abbreviated = renderNumber(num);
    const exact = num.toLocaleString();

    // For charts, return the abbreviated version but include exact in a data attribute or similar
    if (num >= 10000) {
        return `${abbreviated} (${exact})`;
    }

    return abbreviated;
}

export function renderQuotaWithPrompt(quota, digits) {
    let displayInCurrency = localStorage.getItem('display_in_currency');
    displayInCurrency = displayInCurrency === 'true';
    if (displayInCurrency) {
        return `（等价金额：${renderQuota(quota, digits)}）`;
    }
    return '';
}

export function downloadTextAsFile(text, filename) {
    let blob = new Blob([text], {type: 'text/plain;charset=utf-8'});
    let url = URL.createObjectURL(blob);
    let a = document.createElement('a');
    a.href = url;
    a.download = filename;
    a.click();
}

export function removeTrailingSlash(url) {
    if (url.endsWith('/')) {
        return url.slice(0, -1);
    } else {
        return url;
    }
}

let channelModels = undefined;

export async function loadChannelModels() {
    const res = await API.get('/api/models');
    const {success, data} = res.data;
    if (!success) {
        return;
    }
    channelModels = data;
    localStorage.setItem('channel_models', JSON.stringify(data));
}

export function getChannelModels(type) {
    if (channelModels !== undefined && type in channelModels) {
        return channelModels[type];
    }
    let models = localStorage.getItem('channel_models');
    if (!models) {
        return [];
    }
    channelModels = JSON.parse(models);
    if (type in channelModels) {
        return channelModels[type];
    }
    return [];
}

export function copy(text, name = '') {
    if (navigator.clipboard && navigator.clipboard.writeText) {
        navigator.clipboard.writeText(text).then(() => {
            showNotice(`复制${name}成功！`, true);
        }, () => {
            text = `复制${name}失败，请手动复制：<br /><br />${text}`;
            enqueueSnackbar(<SnackbarHTMLContent htmlContent={text}/>, getSnackbarOptions('COPY'));
        });
    } else {
        const textArea = document.createElement("textarea");
        textArea.value = text;
        document.body.appendChild(textArea);
        textArea.select();
        try {
            document.execCommand('copy');
            showNotice(`复制${name}成功！`, true);
        } catch (err) {
            text = `复制${name}失败，请手动复制：<br /><br />${text}`;
            enqueueSnackbar(<SnackbarHTMLContent htmlContent={text}/>, getSnackbarOptions('COPY'));
        }
        document.body.removeChild(textArea);
    }
}
