package app

import (
	"nable.gin/app/models"
	"nable.gin/libraries/db"
)

// 初始化表
func Migrations()  {

	db := db.GetMysql()
	db.AutoMigrate(
		&models.User{},//用户表
		&models.Node{},//节点权限表
		&models.Role{},//角色表
		&models.RoleNode{},//角色节点中间表
	)

	//user := models.User{
	//	TrueName: "koko",
	//	UserName: "koko",
	//	Password: "123456",
	//	Email: "495105@qq.com",
	//	Mobile: "13501891198",
	//	IsActive: 1,
	//}
	//db.Create(&user)


	//插入测试数据
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(1, 1, 'admin', '$2y$10$TWZWlsER73KytXsWi8aa/uRWqTc/jAk24vd4GjYUyeFsqvwHVmrBy', '朱君', 'vel_illo@hotmail.com', '+9211329517036', '先生', '/static/img/person_1.jpg', 1, '', '2020-02-13 18:21:51', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(2, 1, 'manager', '$2y$10$DIrIkrLSFsmTlAFyQIF6Iu/ONpkoHcE2Be1HtGf40TzFGrQg5Y3LW', '彭洁', 'inventore_voluptatem@yahoo.com', '+3836087328871', '先生', '/static/img/person_2.jpg', 1, '', '2020-02-13 18:21:51', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(3, 3, 'iure26', '$2y$10$3Uk7ihat8t8wJr03iGF8tux.S.fHWog0Qm1WZCgKCrZzo0M6VLOry', '杜兰英', 'rvoluptate@sina.com', '+3414160489938', '女士', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:51', '2020-02-13 18:21:51', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(4, 3, 'voluptas67', '$2y$10$EQBTKhzqw30Q4HTZkFgoueLUZ44LsvV4bD.z/hd4v1DkgjdUl8o1K', '薄玉珍', 'et_voluptatem@126.com', '+6028096811296', '女士', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:51', '2020-02-13 18:21:51', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(5, 3, 'voluptatem19', '$2y$10$2UqZf.1pgk2SagwHXNI8DeHNWWp7QhLFs4H5WnsNy4.ttQ6FMV3I2', '时伦', 'unde.voluptatibus@sohu.com', '+5686570505155', '女士', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:51', '2020-02-13 18:21:51', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(6, 3, 'uveniam', '$2y$10$g83.XXqdlOSAlmOeQxnZxe9Us6CxOKrN.WUELENfNu25xekLQByL.', '葛楠', 'id25@qq.com', '+4078728256829', '先生', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:51', '2020-02-13 18:21:51', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(7, 3, 'id.id', '$2y$10$ZtniZcEFKhfTWCay55MmOeJt5RLlrHOeDb3YBZm/u28ITLqyO7EYW', '时平', 'sapiente34@sohu.com', '+9558929390847', '女士', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(8, 3, 'odit_qui', '$2y$10$bDZcOYKSjS7u8mvZZ7pKUeJ.J/VnWBMPV2od3mEQ9npaqiYoAUaQK', '文洪', 'ea_ut@sohu.com', '+7516062191635', '女士', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(9, 3, 'ipsam_omnis', '$2y$10$sDcdPAMMt38o1kGPmYKeVekCDZUtun.WYRWCjTm0gWFqKsv8IZVUC', '饶超', 'wrepudiandae@hotmail.com', '+5035531299386', '先生', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(10, 3, 'sint_velit', '$2y$10$ArbC3l.rH9BaSvfaLLzkcuF/ENWiAGHuUCIt4QvzucVutTIQ3WJTq', '殷晶', 'libero.sapiente@126.com', '+4771361095886', '女士', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(11, 3, 'quia_ducimus', '$2y$10$KAe66/bnYfC5l3b1oxiu2ed1iukfJ6vFXLdpmaeK3FIyv69FEaggG', '娄嘉', 'omnis_aut@sohu.com', '+3433548185632', '女士', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(12, 3, 'dvoluptates', '$2y$10$Cw/XTg8R6j7UFL8XsztpVe2g0P2seWsxwFOkM.tVH.RYa6Z1GIvFq', '盖鑫', 'quis_est@gmail.com', '+1105784299695', '先生', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(13, 3, 'wmaiores', '$2y$10$IDrhEyHvpwFZDVJsJW7/WeO3.DCRKxfyrTZ7qJXL61I5g1bgzT.ZS', '洪智敏', 'eligendi.porro@sohu.com', '+8740078821204', '先生', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(14, 3, 'fdolor', '$2y$10$5zVhX5qX2ga08uGeXLySQuL5kqeniO24W7S/KdT6/6WuF.5C15Kjm', '司凤英', 'consequuntur23@gmail.com', '+6987353268450', '女士', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(15, 3, 'itaque97', '$2y$10$/RrRuQLqvq6e2H6hrVhaoOY0Ho3kEdUjEpvaXx6MtTsCN40.iFBtO', '焦欣', 'reprehenderit.expedita@hotmail.com', '+3848277391253', '先生', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(16, 3, 'qui46', '$2y$10$6vNwzjDfGKKOUYvjEonC.OOXoNv/jyceWYTkGhOIL0JEZ0m0JY2ne', '涂秀云', 'et17@163.com', '+1752617963181', '先生', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(17, 3, 'zvoluptate', '$2y$10$/q9eY8bX./V1EHRe98H6lu07PxppFENbyPIEUHjy3C4i7eqg9EPcG', '阮琴', 'unde_dolores@126.com', '+2915502079068', '先生', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(18, 3, 'quibusdam.id', '$2y$10$zxSqEHJob8Y7QXcIpkgKSepJuADL0AZnz0bJ2l33rkoRzOqz6Ce9m', '冷雪', 'est.tempore@hotmail.com', '+2493885009322', '先生', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(19, 3, 'fdolores', '$2y$10$n6oV7oJAFAL7fYNPSFL.nue5L6Whp9zmywxSzheTwj1fJxtTAQDvq', '凌军', 'bvoluptates@hotmail.com', '+7489423977199', '女士', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(20, 3, 'veritatis.officiis', '$2y$10$9Nfv5c7hbxbXhZigFHM.KO4oXwpS8AhbCCd70Cz7r3IPYN0SUq5Ee', '都明霞', 'quaerat.quidem@qq.com', '+1567767960051', '先生', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(21, 3, 'at.doloribus', '$2y$10$Uu/JXxgayviCYRq1Jg28eu9Ql2sOH5FXk.HjI1nEnjuoFHhWbTDvO', '邬桂芬', 'et.quae@gmail.com', '+9749571685665', '女士', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(22, 3, 'dolor88', '$2y$10$niqUnVRDzLhr3OPX9xAYmekf0pCQjl4FIgqAPlmP72lBtb1wfiAXG', '常琴', 'doloremque25@hotmail.com', '+9454649435558', '女士', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(23, 3, 'ut_occaecati', '$2y$10$4YkQfPCugGchpYckT2Yije1lZBe2L/.jC.VQwkcPCDm1icCQ2EHJa', '岳坤', 'velit.perferendis@126.com', '+7641048010116', '女士', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(24, 3, 'at_illo', '$2y$10$QqupQ//92UysGXlGcciZWeDh706aUDWrsHT8gtxGUv8SSMpe5M2/a', '倪依琳', 'vero_non@yahoo.com', '+9670795218473', '女士', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(25, 3, 'repellendus.omnis', '$2y$10$vHLeSQjVPziIaIh/kisKW.4YuCwpmxOrgfWGfJZlQ70LdazghRFyu', '姬新华', 'et_natus@163.com', '+5619637084567', '女士', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(26, 3, 'non.molestiae', '$2y$10$hFEhdEy69rYAU3QWcx.w7er2U9iF6OUbAH/ydOqlxOkkiW.SJuecm', '巫欢', 'zautem@126.com', '+7216497090096', '先生', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(27, 3, 'wrepellat', '$2y$10$ZRfqB5zUo0TPYmBujklFo.qBGAgqp2/RHFS2t.jQdzivMAEQBxEUS', '郑玉珍', 'doloribus55@sina.com', '+5219837963951', '女士', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(28, 3, 'rqui', '$2y$10$9t4AQXb88nw9AOAodVg4ZOtlL83AnsZr0WP4.9uptY1c36frvTy2e', '郑新华', 'voluptas_rerum@sohu.com', '+4774840299857', '先生', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")
	db.Exec("INSERT INTO `na_users` (`id`, `role_id`, `username`, `password`, `truename`, `email`, `mobile`, `sex`, `avatar`, `is_active`, `last_ip`, `created_at`, `updated_at`, `deleted_at`) VALUES(29, 3, 'numquam.neque', '$2y$10$YNiZvIffj2u4YOXxJ3GMvOXd8Yi62gCgH4tUQduFU5TW4N0juygLK', '殷文彬', 'rquis@qq.com', '+8322980170704', '女士', '/static/img/default.jpg', 1, '', '2020-02-13 18:21:52', '2020-02-13 18:21:52', NULL);")

	db.Exec("INSERT INTO `na_roles` (`id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES(1, '超级管理员', '2020-02-13 18:21:52', NULL, NULL);")
	db.Exec("INSERT INTO `na_roles` (`id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES(2, '管理员', '2020-02-13 18:21:52', NULL, NULL);")
	db.Exec("INSERT INTO `na_roles` (`id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES(3, '普通用户', '2020-02-13 18:21:52', NULL, NULL);")

	db.Exec("INSERT INTO `na_nodes` (`id`, `name`, `icon`, `url`, `route_name`, `pid`, `is_type`, `reorder`, `created_at`, `updated_at`, `deleted_at`) VALUES(1, 'Dashboards', NULL, NULL, NULL, 0, '0', 0, '2020-02-11 21:55:40', '2020-02-11 21:55:40', NULL);")
	db.Exec("INSERT INTO `na_nodes` (`id`, `name`, `icon`, `url`, `route_name`, `pid`, `is_type`, `reorder`, `created_at`, `updated_at`, `deleted_at`) VALUES(2, '分析页', NULL, NULL, NULL, 1, '0', 0, '2020-02-11 21:56:08', '2020-02-11 21:56:08', NULL);")
	db.Exec("INSERT INTO `na_nodes` (`id`, `name`, `icon`, `url`, `route_name`, `pid`, `is_type`, `reorder`, `created_at`, `updated_at`, `deleted_at`) VALUES(3, '站点设置', NULL, NULL, NULL, 1, '0', 0, '2020-02-11 21:56:33', '2020-02-11 21:56:33', NULL);")
	db.Exec("INSERT INTO `na_nodes` (`id`, `name`, `icon`, `url`, `route_name`, `pid`, `is_type`, `reorder`, `created_at`, `updated_at`, `deleted_at`) VALUES(4, '用户管理', NULL, NULL, NULL, 0, '0', 0, '2020-02-11 21:56:48', '2020-02-11 21:56:48', NULL);")
	db.Exec("INSERT INTO `na_nodes` (`id`, `name`, `icon`, `url`, `route_name`, `pid`, `is_type`, `reorder`, `created_at`, `updated_at`, `deleted_at`) VALUES(5, '用户列表', NULL, NULL, NULL, 4, '0', 0, '2020-02-11 21:57:06', '2020-02-11 21:57:06', NULL);")
	db.Exec("INSERT INTO `na_nodes` (`id`, `name`, `icon`, `url`, `route_name`, `pid`, `is_type`, `reorder`, `created_at`, `updated_at`, `deleted_at`) VALUES(6, '权限管理', NULL, NULL, NULL, 0, '0', 0, '2020-02-11 21:57:38', '2020-02-11 21:57:38', NULL);")
	db.Exec("INSERT INTO `na_nodes` (`id`, `name`, `icon`, `url`, `route_name`, `pid`, `is_type`, `reorder`, `created_at`, `updated_at`, `deleted_at`) VALUES(7, '角色管理', NULL, NULL, NULL, 6, '0', 0, '2020-02-11 21:57:56', '2020-02-11 21:57:56', NULL);")
	db.Exec("INSERT INTO `na_nodes` (`id`, `name`, `icon`, `url`, `route_name`, `pid`, `is_type`, `reorder`, `created_at`, `updated_at`, `deleted_at`) VALUES(8, '权限管理', NULL, NULL, NULL, 6, '0', 0, '2020-02-11 21:58:20', '2020-02-11 21:58:20', NULL);")
	db.Exec("INSERT INTO `na_nodes` (`id`, `name`, `icon`, `url`, `route_name`, `pid`, `is_type`, `reorder`, `created_at`, `updated_at`, `deleted_at`) VALUES(9, '管理员列表', NULL, NULL, NULL, 6, '0', 0, '2020-02-11 21:58:37', '2020-02-11 21:58:37', NULL);")






}

