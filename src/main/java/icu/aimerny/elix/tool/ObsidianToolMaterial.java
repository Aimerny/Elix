package icu.aimerny.elix.tool;

import icu.aimerny.elix.registry.ModItem;
import net.minecraft.item.*;
import net.minecraft.recipe.Ingredient;

public class ObsidianToolMaterial implements ToolMaterial {

    private static final ObsidianToolMaterial INSTANCE = new ObsidianToolMaterial();
    public static final ToolItem OBSIDIAN_PICKAXE = new PickaxeItem(INSTANCE, 5, 2.0F,new Item.Settings());
    public static final ToolItem OBSIDIAN_AXE = new AxeItem(INSTANCE, 12, 0.6F,new Item.Settings());
    public static final ToolItem OBSIDIAN_HOE = new HoeItem(INSTANCE, 5, 1.0F,new Item.Settings());
    public static final ToolItem OBSIDIAN_SHOVEL = new ShovelItem(INSTANCE, 5, 5.0F,new Item.Settings());
    public static final ToolItem OBSIDIAN_SWORD = new SwordItem(INSTANCE, 9, 2.0F,new Item.Settings());

    @Override
    public int getDurability() {
        return 3500;
    }

    @Override
    public float getMiningSpeedMultiplier() {
        return 16.0F;
    }

    @Override
    public float getAttackDamage() {
        return 5.0F;
    }

    @Override
    public int getMiningLevel() {
        return 3;
    }

    @Override
    public int getEnchantability() {
        return 50;
    }

    @Override
    public Ingredient getRepairIngredient() {
        return Ingredient.ofItems(ModItem.OBSIDIAN_INGOT);
    }

    public static void init() {
        ModItem.register("obsidian_pickaxe", OBSIDIAN_PICKAXE);
        ModItem.register("obsidian_axe", OBSIDIAN_AXE);
        ModItem.register("obsidian_hoe", OBSIDIAN_HOE);
        ModItem.register("obsidian_sword", OBSIDIAN_SWORD);
        ModItem.register("obsidian_shovel", OBSIDIAN_SHOVEL);
    }
}
