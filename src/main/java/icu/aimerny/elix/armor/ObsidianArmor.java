package icu.aimerny.elix.armor;

import icu.aimerny.elix.consts.IdConst;
import icu.aimerny.elix.registry.ModItem;
import net.minecraft.entity.EquipmentSlot;
import net.minecraft.item.ArmorItem;
import net.minecraft.item.Item;
import net.minecraft.item.ItemStack;
import net.minecraft.recipe.Ingredient;
import net.minecraft.sound.SoundEvent;
import net.minecraft.sound.SoundEvents;

import java.util.List;

public class ObsidianArmor extends ElixArmor{


    public static final ObsidianArmor OBSIDIAN_ARMOR_MATERIAL = new ObsidianArmor();
    public static final Item OBSIDIAN_MATERIAL_HELMET = new ArmorItem(OBSIDIAN_ARMOR_MATERIAL, ArmorItem.Type.HELMET, new Item.Settings());
    public static final Item OBSIDIAN_MATERIAL_CHESTPLATE = new ArmorItem(OBSIDIAN_ARMOR_MATERIAL, ArmorItem.Type.CHESTPLATE, new Item.Settings());
    public static final Item OBSIDIAN_MATERIAL_LEGGINGS = new ArmorItem(OBSIDIAN_ARMOR_MATERIAL, ArmorItem.Type.LEGGINGS, new Item.Settings());
    public static final Item OBSIDIAN_MATERIAL_BOOTS = new ArmorItem(OBSIDIAN_ARMOR_MATERIAL, ArmorItem.Type.BOOTS, new Item.Settings());



    @Override
    public int getDurability(ArmorItem.Type type) {
        return BASE_DURABILITY[type.getEquipmentSlot().getEntitySlotId()];
    }

    @Override
    public int getProtection(ArmorItem.Type type) {
        return PROTECTION_VALUES[type.getEquipmentSlot().getEntitySlotId()];
    }

    @Override
    public int getEnchantability() {
        return 2;
    }

    @Override
    public SoundEvent getEquipSound() {
        return SoundEvents.ITEM_ARMOR_EQUIP_CHAIN;
    }

    @Override
    public Ingredient getRepairIngredient() {
        return Ingredient.ofItems(ModItem.OBSIDIAN_INGOT);
    }

    @Override
    public String getName() {
        return "obsidian";
    }

    @Override
    public float getToughness() {
        return 1;
    }

    @Override
    public float getKnockbackResistance() {
        return 0;
    }

    public static void init(){
        ModItem.register(IdConst.OBSIDIAN_MATERIAL_HELMET, OBSIDIAN_MATERIAL_HELMET);
        ModItem.register(IdConst.OBSIDIAN_MATERIAL_CHESTPLATE, OBSIDIAN_MATERIAL_CHESTPLATE);
        ModItem.register(IdConst.OBSIDIAN_MATERIAL_LEGGINGS, OBSIDIAN_MATERIAL_LEGGINGS);
        ModItem.register(IdConst.OBSIDIAN_MATERIAL_BOOTS, OBSIDIAN_MATERIAL_BOOTS);
    }
}
