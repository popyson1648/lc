"use strict";
const fs = require('fs');
const path = require('path');

function getProblemDirPath() {
    // カレントディレクトリの config-path.json の "leetcodeDirPath" の値を取得
    const configPath = path.join(__dirname, 'config-path.json');
    const configData = JSON.parse(fs.readFileSync(configPath, 'utf8'));
    const leetcodeDirPath = configData.leetcodeDirPath;

    // leetcodeDirPath ディレクトリ内の config.json の中の "problemDirPath" の値を返す
    const leetcodeConfigPath = path.join(leetcodeDirPath, 'config.json');
    const leetcodeConfigData = JSON.parse(fs.readFileSync(leetcodeConfigPath, 'utf8'));
    return leetcodeConfigData.problemDirPath;
}

module.exports = {
    getProblemDirPath
};
