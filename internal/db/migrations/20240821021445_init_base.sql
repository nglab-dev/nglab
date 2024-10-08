-- +goose Up
-- +goose StatementBegin

-- 字段类型表
CREATE TABLE `base_dict_type` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` TEXT NOT NULL, -- 类型名称
    `remark` TEXT DEFAULT NULL, -- 备注
    `sort` INTEGER DEFAULT 0, -- 排序
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    `created_by` INTEGER NOT NULL DEFAULT 0, -- 创建人
    `updated_by` INTEGER NOT NULL DEFAULT 0 -- 更新人
);

INSERT INTO "base_dict_type" ("id", "name") VALUES (1, '样本类型');
INSERT INTO "base_dict_type" ("id", "name") VALUES (2, '实验方法');
INSERT INTO "base_dict_type" ("id", "name") VALUES (3, '结果单位');
INSERT INTO "base_dict_type" ("id", "name") VALUES (4, '标本性状');
INSERT INTO "base_dict_type" ("id", "name") VALUES (5, '禁止打印原因');

-- 字典表
CREATE TABLE `base_dict` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `type` INTEGER NOT NULL, --字典类型，1 样本类型 2 实验方法 3 结果单位 4 标本性状 5 禁止打印原因
    `name` TEXT NOT NULL, --字典名称
    `alias` TEXT DEFAULT NULL, --别名
    `remark` TEXT DEFAULT NULL, --备注
    `sort` INTEGER DEFAULT 0, -- 排序
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP, -- 更新时间
    `created_by` INTEGER NOT NULL DEFAULT 0, -- 创建人
    `updated_by` INTEGER NOT NULL DEFAULT 0 -- 更新人
);

-- 1. 标本类型插入
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '血液', 'Blood', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '血清', 'Serum', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '血浆', 'Plasma', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '尿液', 'Urine', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '24h尿', NULL, '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '粪便', 'Stool', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '全血', 'whole blood', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '动脉血', 'Arterial Blood', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '末梢血', 'Peripheral Blood', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '阴道分泌物', NULL, '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '前列腺液', NULL, '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '脑脊液', NULL, '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '穿刺液', NULL, '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '腹水', 'Abdominal fluid', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '咽拭子', NULL, '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '呕吐物', NULL, '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '阴道拭子', NULL, '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '鼻咽拭子', NULL, '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '精液', 'Semen', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '拭子', 'Swab', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '胃液', 'Gastric fluid', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '痰', 'Sputum', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '脓液', 'Pus', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '胸水', 'Pleural fluid', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '腹水', NULL, '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '关节腔积液', 'Joint fluid', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '心包积液', 'Pericardial fluid', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '腹腔积液', NULL, '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '胆汁', 'Bile', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '脑脊液', 'Cerebrospinal fluid', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '腹腔积液', NULL, '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '前列腺液', 'Prostatic fluid', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '体液', 'Fluid', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '羊水', 'Amniotic fluid', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '分泌物', 'Secretion', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '透析液', 'Dialysis fluid', '常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '胃内容物', NULL, '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '胸腔积液', NULL, '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '支气管肺泡灌洗液', 'Broncho-alveolar lavage', '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '引流液', 'Shunt fluid', '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '导尿管尿', NULL, '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '肾穿刺尿', NULL, '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '前段尿', NULL, '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '伤口渗液', NULL, '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '乳汁', 'Breast milk', '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '灌洗液', NULL, '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '积液', NULL, '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '腹腔引流液', NULL, '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '渗出液', NULL, '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '膝关节液', 'Knee fluid', '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '眼分泌物', NULL, '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '脓液', NULL, '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '伤口拭子', NULL, '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '引流物', NULL, '一般');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '食道粘膜白苔', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '囊中穿刺液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '脓疱液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '腹腔渗液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '关节拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '外科手术伤口拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '尿道拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '睾丸穿刺液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '水泡液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '疱疹液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '结膜囊分泌物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '口腔拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '渗出物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '导管血', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '中耳拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '皮下积液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '引流管液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '外耳拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '切口分泌物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '脐部拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '肾积水', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '积血', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '右乳突腔脓性物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '阴囊内脓液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '睾丸内脓液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '刀口下积液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '肾周积水', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '肾周积液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '口内分泌物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '伤口分泌物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '上额窦内分泌物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '疱液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '尿道分泌物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '右耳乳突腔内脓性物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '组织穿刺物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '创面分泌物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '腹腔内容物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '引流血液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '角膜分泌物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '支气管分泌物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '食道附着物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '耳部标本', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '皮肤拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '肺穿刺液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '窦道拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '无菌体液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '直肠拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '手部拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '结膜拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '生殖器拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '男性生殖器拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '女性生殖器拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '足拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '疖', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '乳房穿刺液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '下呼吸道标本', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '上呼吸道标本', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '清洁中段尿', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '筛选MRSA标本', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '筛选VRE标本', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '科研标本', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '实验室标本', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '呼吸道标本', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '溃疡标本', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '胸腹水', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '渗液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '左胸水', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '右胸水', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '宫腔分泌物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '穿刺液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '宫颈脱落细胞', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '食管粘膜白苔', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '压疮拭子', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '气管吸出物', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '脑室引流液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '骨髓穿刺液', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '不明标本', NULL, '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '水', 'Water', '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '环境', 'Environmental', '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '胎尿、胎粪', 'Meconium', '不常用');
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (1, '其它', 'Other', '不常用');


-- 2. 实验方法插入
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '化学发光法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '速率法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '酶法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '酶联免疫法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '胶体金法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '金标法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '金标免疫法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '荧光PCR法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '干化学法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '免疫比浊法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '高效液相法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '手工法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '散射比浊法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '免疫凝集法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '凝集法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '阴抗法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '激光散色法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '荧光染色法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '流式细胞术', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '双缩脲法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '胆固醇过氧化酶法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '甘油磷酸氧化酶法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '尿酸酶终点法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '尿素酶-GLDH法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '氧化酶法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '钒酸盐氧化法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '溴甲酚绿', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, 'ADA LIQUID REAGENT', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '双缩脲比色法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '速率法（AMP缓冲）', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, 'Calmagite比色法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, 'Dot-ELISA', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, 'MTB比色法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, 'RNA捕获探针法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '底物酶法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '干式免疫荧光定量法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '化学修饰酶法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '己糖激酶法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '离子选择电极法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '磷钼酸比色法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '酶比色法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '酶速率法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '免疫层析法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '目测', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '脲酶速率法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '凝固法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '速率法L-P', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '速率法酶', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '微量蛋白染料法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '溴甲酚绿法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '一点法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '两点法', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (2, '中点法', NULL, NULL);

-- 3. 结果单位插入
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '%', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '/Hp', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '/HPF', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '/LPF', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '°C', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '10^12/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '10^6/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '10^6/mL', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '10^9/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '200/s', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '5ug/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'AN/ml', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'AU/ml', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'CELL/u', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'Cell/μl', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'cfu/cm2', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'cfu/ml', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'cfu/件', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'ch', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'CO/S', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'COI', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'copies/ml', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'd/sc', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'deg', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'dyn/cm^2', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'e12/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'e9.SI', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'fL', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'g', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'g/dL', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'g/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'HPF', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'Inh%', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'IU/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'IU/mL', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'KU/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'L/24h', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'L/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'Leu/uL', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mg', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mg/12h', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mg/24h', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mg/8h', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mg/dL', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mg/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mg/mmol', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mg/天', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'min', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mIU/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mIU/mL', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'ml', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'ml/min', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mm', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mm/h', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mm^3', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mmHg', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mmmol', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mmol', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mmol/24h', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mmol/d', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mmol/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mmol/mol', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mPa', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mPa.S', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mPas', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mPs.S', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mS/cm', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'mu/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'ng/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'ng/ml', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'nmmol', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'nmol/l', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'PEIU/mL', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'pg', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'pg/ml', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'pmol/l', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'pomo/l', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'S', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'S/CO', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'Sec', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'U', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'U/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'U/ml', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'ug/dl', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'ug/g', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'ug/l', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'ug/mL', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'uIU/ml', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'uL', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'ummol/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'umol/24h', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'umol/6h', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'umol/L', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'μg/l', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, 'μIU/ml', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '百万/毫升', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '分钟', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '个', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '个/HP', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '个/LP', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '个/ul', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '个/μL', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '级', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '千克', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (3, '天', NULL, NULL);

-- 4. 标本性状插入
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (4, '正常', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (4, '溶血', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (4, '乳糜', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (4, '轻度溶血', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (4, '严重溶血', NULL, NULL);
INSERT INTO "base_dict" ("type", "name", "alias", "remark") VALUES (4, '严重乳糜', NULL, NULL);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `base_dict`;
-- +goose StatementEnd
